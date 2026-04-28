package depsdiff_test

import (
	"fmt"
	"testing"

	depsdiff "github.com/yoanm/go-deps-diff"
	"github.com/yoanm/go-deps-diff/composer"
)

func TestDiff_BasicComparison(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		lockPrevious   []byte
		lockCurrent    []byte
		reqPrevious    []byte
		reqCurrent     []byte
		checkResultsFn func(output *depsdiff.Output) bool
	}{
		{
			name: "identical locks",
			lockPrevious: []byte(`{
				"packages": [{"name": "vendor/pkg", "version": "1.0.0"}]
			}`),
			lockCurrent: []byte(`{
				"packages": [{"name": "vendor/pkg", "version": "1.0.0"}]
			}`),
			reqPrevious: []byte(`{"require": {"vendor/pkg": "^1.0"}}`),
			reqCurrent:  []byte(`{"require": {"vendor/pkg": "^1.0"}}`),
			checkResultsFn: func(out *depsdiff.Output) bool {
				return len(out.Changes) == 0
			},
		},
		{
			name:         "added package",
			lockPrevious: []byte(`{"packages": []}`),
			lockCurrent: []byte(`{
				"packages": [{"name": "vendor/new", "version": "1.0.0", "source": {"reference": "abc"}}]
			}`),
			reqPrevious: []byte(`{"require": {}}`),
			reqCurrent:  []byte(`{"require": {"vendor/new": "^1.0"}}`),
			checkResultsFn: func(out *depsdiff.Output) bool {
				if len(out.Changes) != 1 {
					return false
				}

				return out.Changes["vendor/new"].Operation.Name == depsdiff.AdditionOperation
			},
		},
		{
			name: "removed package",
			lockPrevious: []byte(`{
				"packages": [{"name": "vendor/old", "version": "1.0.0", "source": {"reference": "abc"}}]
			}`),
			lockCurrent: []byte(`{"packages": []}`),
			reqPrevious: []byte(`{"require": {"vendor/old": "^1.0"}}`),
			reqCurrent:  []byte(`{"require": {}}`),
			checkResultsFn: func(out *depsdiff.Output) bool {
				if len(out.Changes) != 1 {
					return false
				}

				return out.Changes["vendor/old"].Operation.Name == depsdiff.RemovalOperation
			},
		},
		{
			name: "updated package",
			lockPrevious: []byte(`{
				"packages": [{"name": "vendor/pkg", "version": "1.0.0", "source": {"reference": "abc"}}]
			}`),
			lockCurrent: []byte(`{
				"packages": [{"name": "vendor/pkg", "version": "2.0.0", "source": {"reference": "def"}}]
			}`),
			reqPrevious: []byte(`{"require": {"vendor/pkg": "*"}}`),
			reqCurrent:  []byte(`{"require": {"vendor/pkg": "*"}}`),
			checkResultsFn: func(out *depsdiff.Output) bool {
				if len(out.Changes) != 1 {
					return false
				}

				change := out.Changes["vendor/pkg"]

				return change.Operation.Name == depsdiff.UpgradeOperation &&
					change.Operation.SemverType == "MAJOR"
			},
		},
	}

	for _, testData := range tests {
		t.Run(testData.name, func(t *testing.T) {
			t.Parallel()

			out, err := buildMapFromFixtures(testData)
			if err != nil {
				t.Fatal(err)
			}

			if !testData.checkResultsFn(out) {
				t.Errorf("Diff() result check failed")
			}
		})
	}
}

func buildMapFromFixtures(testData struct {
	name           string
	lockPrevious   []byte
	lockCurrent    []byte
	reqPrevious    []byte
	reqCurrent     []byte
	checkResultsFn func(output *depsdiff.Output) bool
},
) (*depsdiff.Output, error) {
	previousMap, err := composer.BuildMapFromBytes(testData.reqPrevious, testData.lockPrevious)
	if err != nil {
		return nil, fmt.Errorf("parsing previous data: %w", err)
	}

	currentMap, err := composer.BuildMapFromBytes(testData.reqCurrent, testData.lockCurrent)
	if err != nil {
		return nil, fmt.Errorf("parsing current data: %w", err)
	}

	out, err := depsdiff.Diff(previousMap, currentMap)
	if err != nil {
		return nil, fmt.Errorf("error during diff process: %w", err)
	}

	return out, nil
}

func TestDiff_IsRootRequirement(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		lockPrevious []byte
		lockCurrent  []byte
		reqPrevious  []byte
		reqCurrent   []byte
		checkFn      func(output *depsdiff.Output) bool
	}{
		{
			name:         "with composer.json require",
			lockPrevious: []byte(`{"packages": []}`),
			lockCurrent: []byte(`{
				"packages": [{"name": "vendor/pkg", "version": "1.0.0", "source": {"reference": "abc"}}]
			}`),
			reqPrevious: []byte(`{"require": {}}`),
			reqCurrent:  []byte(`{"require": {"vendor/pkg": "^1.0"}}`),
			checkFn: func(out *depsdiff.Output) bool {
				if len(out.Changes) == 0 {
					return false
				}

				return out.Changes["vendor/pkg"].Package.IsRootRequirement() == true
			},
		},
		{
			name:         "with empty composer.json",
			lockPrevious: []byte(`{"packages": []}`),
			lockCurrent: []byte(`{
				"packages": [{"name": "vendor/pkg", "version": "1.0.0", "source": {"reference": "abc"}}]
			}`),
			reqPrevious: []byte(`{"require": {}}`),
			reqCurrent:  []byte(`{"require": {}}`),
			checkFn: func(out *depsdiff.Output) bool {
				if len(out.Changes) == 0 {
					return false
				}

				return out.Changes["vendor/pkg"].Package.IsRootRequirement() == false
			},
		},
		{
			name:         "with composer.json require-dev",
			lockPrevious: []byte(`{"packages": []}`),
			lockCurrent: []byte(`{
				"packages-dev": [{"name": "vendor/test", "version": "1.0.0", "source": {"reference": "abc"}}]
			}`),
			reqPrevious: []byte(`{"require": {}}`),
			reqCurrent:  []byte(`{"require-dev": {"vendor/test": "^1.0"}}`),
			checkFn: func(out *depsdiff.Output) bool {
				if len(out.Changes) == 0 {
					return false
				}

				return out.Changes["vendor/test"].Package.IsRootDevRequirement() == true
			},
		},
	}

	for _, testData := range tests {
		t.Run(testData.name, func(t *testing.T) {
			t.Parallel()

			previousMap, err := composer.BuildMapFromBytes(testData.reqPrevious, testData.lockPrevious)
			if err != nil {
				t.Error(fmt.Errorf("parsing previous data: %w", err))

				return
			}

			currentMap, err := composer.BuildMapFromBytes(testData.reqCurrent, testData.lockCurrent)
			if err != nil {
				t.Error(fmt.Errorf("parsing current data: %w", err))

				return
			}

			out, err := depsdiff.Diff(previousMap, currentMap)
			if err != nil {
				t.Errorf("Diff() error = %v", err)

				return
			}

			if testData.checkFn != nil && !testData.checkFn(out) {
				t.Errorf("Diff() result check failed")
			}
		})
	}
}

func TestDiff_DevPackages(t *testing.T) {
	t.Parallel()

	lockDataPrevious := []byte(`{
		"packages": [{"name": "vendor/lib", "version": "1.0.0", "source": {"reference": "abc"}}],
		"packages-dev": [{"name": "vendor/test", "version": "1.0.0", "source": {"reference": "def"}}]
	}`)
	lockDataCurrent := []byte(`{
		"packages": [{"name": "vendor/lib", "version": "1.0.0", "source": {"reference": "abc"}}],
		"packages-dev": [{"name": "vendor/test", "version": "2.0.0", "source": {"reference": "def"}}]
	}`)
	reqData := []byte(`{"require": {"vendor/lib": "^1.0"}, "require-dev": {"vendor/test": "^2.0"}}`)

	previousMap, err := composer.BuildMapFromBytes(reqData, lockDataPrevious)
	if err != nil {
		t.Error(fmt.Errorf("parsing previous data: %w", err))

		return
	}

	currentMap, err := composer.BuildMapFromBytes(reqData, lockDataCurrent)
	if err != nil {
		t.Error(fmt.Errorf("parsing current data: %w", err))

		return
	}

	out, err := depsdiff.Diff(previousMap, currentMap)
	if err != nil {
		t.Errorf("Diff() error = %v", err)

		return
	}

	if len(out.Changes) != 1 {
		t.Errorf("Diff() got %d packages, want 1", len(out.Changes))

		return
	}

	if _, exist := out.Changes["vendor/test"]; !exist {
		t.Error("Diff() package vendor/test must exist")
	}
}
