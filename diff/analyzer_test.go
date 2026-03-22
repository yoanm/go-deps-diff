package diff

import (
	"testing"
)

func TestDiff_BasicComparison(t *testing.T) {
	tests := []struct {
		name             string
		lockPrevious     []byte
		lockCurrent      []byte
		reqPrevious      []byte
		reqCurrent       []byte
		wantError        bool
		wantPackageCount int
		checkResultsFn   func(*Output) bool
	}{
		{
			name: "identical locks",
			lockPrevious: []byte(`{
				"packages": [{"name": "vendor/pkg", "version": "1.0.0"}]
			}`),
			lockCurrent: []byte(`{
				"packages": [{"name": "vendor/pkg", "version": "1.0.0"}]
			}`),
			wantError:        false,
			wantPackageCount: 0,
		},
		{
			name:         "added package",
			lockPrevious: []byte(`{"packages": []}`),
			lockCurrent: []byte(`{
				"packages": [{"name": "vendor/new", "version": "1.0.0", "source": {"reference": "abc"}}]
			}`),
			wantError:        false,
			wantPackageCount: 1,
			checkResultsFn: func(out *Output) bool {
				if len(out.Packages) != 1 {
					return false
				}
				return out.Packages[0].Name == "vendor/new" &&
					out.Packages[0].Update.Type == "ADDED"
			},
		},
		{
			name: "removed package",
			lockPrevious: []byte(`{
				"packages": [{"name": "vendor/old", "version": "1.0.0", "source": {"reference": "abc"}}]
			}`),
			lockCurrent:      []byte(`{"packages": []}`),
			wantError:        false,
			wantPackageCount: 1,
			checkResultsFn: func(out *Output) bool {
				if len(out.Packages) != 1 {
					return false
				}
				return out.Packages[0].Name == "vendor/old" &&
					out.Packages[0].Update.Type == "REMOVED"
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
			wantError:        false,
			wantPackageCount: 1,
			checkResultsFn: func(out *Output) bool {
				if len(out.Packages) != 1 {
					return false
				}
				pkg := out.Packages[0]
				return pkg.Name == "vendor/pkg" &&
					pkg.Update.Type == "UPDATED" &&
					pkg.Update.SubType == "MAJOR" &&
					pkg.Update.Direction == "UP"
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := Diff(tt.lockPrevious, tt.lockCurrent, tt.reqPrevious, tt.reqCurrent)
			if (err != nil) != tt.wantError {
				t.Errorf("Diff() error = %v, wantError %v", err, tt.wantError)
				return
			}
			if !tt.wantError {
				if len(out.Packages) != tt.wantPackageCount {
					t.Errorf("Diff() got %d packages, want %d", len(out.Packages), tt.wantPackageCount)
				}
				if tt.checkResultsFn != nil && !tt.checkResultsFn(out) {
					t.Errorf("Diff() result check failed")
				}
			}
		})
	}
}

func TestDiff_MutualDependency(t *testing.T) {
	lock := []byte(`{"packages": []}`)

	tests := []struct {
		name        string
		reqPrevious []byte
		reqCurrent  []byte
		wantError   bool
	}{
		{
			name:        "both json provided",
			reqPrevious: []byte(`{}`),
			reqCurrent:  []byte(`{}`),
			wantError:   false,
		},
		{
			name:        "both json nil",
			reqPrevious: nil,
			reqCurrent:  nil,
			wantError:   false,
		},
		{
			name:        "only reqPrevious provided",
			reqPrevious: []byte(`{}`),
			reqCurrent:  nil,
			wantError:   true,
		},
		{
			name:        "only reqCurrent provided",
			reqPrevious: nil,
			reqCurrent:  []byte(`{}`),
			wantError:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Diff(lock, lock, tt.reqPrevious, tt.reqCurrent)
			if (err != nil) != tt.wantError {
				t.Errorf("Diff() error = %v, wantError %v", err, tt.wantError)
			}
		})
	}
}

func TestDiff_IsRootRequirement(t *testing.T) {
	tests := []struct {
		name         string
		lockPrevious []byte
		lockCurrent  []byte
		reqPrevious  []byte
		reqCurrent   []byte
		checkFn      func(*Output) bool
	}{
		{
			name:         "with composer.json require",
			lockPrevious: []byte(`{"packages": []}`),
			lockCurrent: []byte(`{
				"packages": [{"name": "vendor/pkg", "version": "1.0.0", "source": {"reference": "abc"}}]
			}`),
			reqPrevious: []byte(`{}`),
			reqCurrent:  []byte(`{"require": {"vendor/pkg": "^1.0"}}`),
			checkFn: func(out *Output) bool {
				if len(out.Packages) == 0 {
					return false
				}
				return out.Packages[0].IsRootRequirement == true
			},
		},
		{
			name:         "without composer.json",
			lockPrevious: []byte(`{"packages": []}`),
			lockCurrent: []byte(`{
				"packages": [{"name": "vendor/pkg", "version": "1.0.0", "source": {"reference": "abc"}}]
			}`),
			reqPrevious: nil,
			reqCurrent:  nil,
			checkFn: func(out *Output) bool {
				if len(out.Packages) == 0 {
					return false
				}
				return out.Packages[0].IsRootRequirement == false
			},
		},
		{
			name:         "with composer.json require-dev",
			lockPrevious: []byte(`{"packages": []}`),
			lockCurrent: []byte(`{
				"packages-dev": [{"name": "vendor/test", "version": "1.0.0", "source": {"reference": "abc"}}]
			}`),
			reqPrevious: []byte(`{}`),
			reqCurrent:  []byte(`{"require-dev": {"vendor/test": "^1.0"}}`),
			checkFn: func(out *Output) bool {
				if len(out.Packages) == 0 {
					return false
				}
				return out.Packages[0].IsRootDevRequirement == true
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := Diff(tt.lockPrevious, tt.lockCurrent, tt.reqPrevious, tt.reqCurrent)
			if err != nil {
				t.Errorf("Diff() error = %v", err)
				return
			}
			if tt.checkFn != nil && !tt.checkFn(out) {
				t.Errorf("Diff() result check failed")
			}
		})
	}
}

func TestParseSemver(t *testing.T) {
	tests := []struct {
		name    string
		version string
		wantNil bool
		check   func(*PkgVersionTag) bool
	}{
		{
			name:    "simple semver",
			version: "1.2.3",
			wantNil: false,
			check: func(tag *PkgVersionTag) bool {
				return tag.Major == "1" && tag.Minor == "2" && tag.Patch == "3" && tag.Extra == ""
			},
		},
		{
			name:    "with v prefix",
			version: "v2.0.0",
			wantNil: false,
			check: func(tag *PkgVersionTag) bool {
				return tag.Major == "2" && tag.Minor == "0" && tag.Patch == "0"
			},
		},
		{
			name:    "with prerelease",
			version: "1.0.0-beta.1",
			wantNil: false,
			check: func(tag *PkgVersionTag) bool {
				return tag.Major == "1" && tag.Extra == "-beta.1"
			},
		},
		{
			name:    "with build metadata",
			version: "2.5.1+build.1",
			wantNil: false,
			check: func(tag *PkgVersionTag) bool {
				return tag.Extra == "+build.1"
			},
		},
		{
			name:    "invalid semver - commit hash",
			version: "abc123def456",
			wantNil: true,
		},
		{
			name:    "invalid semver - dev-master",
			version: "dev-master",
			wantNil: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tag := parseSemver(tt.version)
			if (tag == nil) != tt.wantNil {
				t.Errorf("parseSemver() = %v, wantNil %v", tag, tt.wantNil)
				return
			}
			if !tt.wantNil && tt.check != nil && !tt.check(tag) {
				t.Errorf("parseSemver() result check failed")
			}
		})
	}
}

func TestDetectUpdate(t *testing.T) {
	tests := []struct {
		name            string
		versionPrevious string
		versionCurrent  string
		wantSubType     string
		wantDirection   string
	}{
		{
			name:            "major version up",
			versionPrevious: "1.0.0",
			versionCurrent:  "2.0.0",
			wantSubType:     "MAJOR",
			wantDirection:   "UP",
		},
		{
			name:            "major version down",
			versionPrevious: "2.0.0",
			versionCurrent:  "1.0.0",
			wantSubType:     "MAJOR",
			wantDirection:   "DOWN",
		},
		{
			name:            "minor version up",
			versionPrevious: "1.0.0",
			versionCurrent:  "1.1.0",
			wantSubType:     "MINOR",
			wantDirection:   "UP",
		},
		{
			name:            "patch version up",
			versionPrevious: "1.0.0",
			versionCurrent:  "1.0.1",
			wantSubType:     "PATCH",
			wantDirection:   "UP",
		},
		{
			name:            "same version",
			versionPrevious: "1.0.0",
			versionCurrent:  "1.0.0",
			wantSubType:     "NONE",
			wantDirection:   "NONE",
		},
		{
			name:            "unparseable version",
			versionPrevious: "abc123",
			versionCurrent:  "def456",
			wantSubType:     "NONE", // @TODO: COPILOT - This must be UNKNOWN ! Update code accordingly
			wantDirection:   "UNKNOWN",
		},
		{
			name:            "one unparseable",
			versionPrevious: "1.0.0",
			versionCurrent:  "dev-master",
			wantSubType:     "NONE", // @TODO: COPILOT - This must be UNKNOWN ! Update code accordingly
			wantDirection:   "UNKNOWN",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := detectUpdate(tt.versionPrevious, tt.versionCurrent)
			if result.SubType != tt.wantSubType {
				t.Errorf("detectUpdate() SubType = %s, want %s", result.SubType, tt.wantSubType)
			}
			if result.Direction != tt.wantDirection {
				t.Errorf("detectUpdate() Direction = %s, want %s", result.Direction, tt.wantDirection)
			}
		})
	}
}

func TestDiff_DevPackages(t *testing.T) {
	lockPrevious := []byte(`{
		"packages": [{"name": "vendor/lib", "version": "1.0.0", "source": {"reference": "abc"}}],
		"packages-dev": [{"name": "vendor/test", "version": "1.0.0", "source": {"reference": "def"}}]
	}`)
	lockCurrent := []byte(`{
		"packages": [{"name": "vendor/lib", "version": "1.0.0", "source": {"reference": "abc"}}],
		"packages-dev": [{"name": "vendor/test", "version": "2.0.0", "source": {"reference": "def"}}]
	}`)

	out, err := Diff(lockPrevious, lockCurrent, nil, nil)
	if err != nil {
		t.Errorf("Diff() error = %v", err)
		return
	}

	if len(out.Packages) != 1 { // @TODO: COPILOT - This should be 2 (there must be 'vendor/lib' and 'vendor/test' packages) ! Update code accordingly
		t.Errorf("Diff() got %d packages, want 1", len(out.Packages))
		return
	}

	if out.Packages[0].Name != "vendor/test" {
		t.Errorf("Diff() package = %s, want vendor/test", out.Packages[0].Name)
	}
}
