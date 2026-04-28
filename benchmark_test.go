package depsdiff_test

import (
	"encoding/json"
	"fmt"
	"testing"

	depsdiff "github.com/yoanm/go-deps-diff"
	"github.com/yoanm/go-deps-diff/composer"
)

func BenchmarkDiff_100Packages(b *testing.B) {
	lockPrevious := generateLockFile(100)
	lockCurrent := generateLockFile(100)

	previousMap, err := composer.BuildMapFromBytes([]byte("{\"require\":{}}"), lockPrevious)
	if err != nil {
		panic(err)
	}

	currentMap, err := composer.BuildMapFromBytes([]byte("{\"require\":{}}"), lockCurrent)
	if err != nil {
		panic(err)
	}

	b.ResetTimer()

	for range b.N {
		if _, err2 := depsdiff.Diff(previousMap, currentMap); err2 != nil {
			b.Fatalf("Diff failed: %v", err2)
		}
	}
}

func BenchmarkDiff_1000Packages(b *testing.B) {
	lockPrevious := generateLockFile(1000)
	lockCurrent := generateLockFile(1000)

	previousMap, err := composer.BuildMapFromBytes([]byte("{\"require\":{}}"), lockPrevious)
	if err != nil {
		panic(err)
	}

	currentMap, err := composer.BuildMapFromBytes([]byte("{\"require\":{}}"), lockCurrent)
	if err != nil {
		panic(err)
	}

	b.ResetTimer()

	for range b.N {
		if _, err2 := depsdiff.Diff(previousMap, currentMap); err2 != nil {
			b.Fatalf("Diff failed: %v", err2)
		}
	}
}

func generateLockFile(count int) []byte {
	type pkg struct {
		Name      string            `json:"name"`
		Version   string            `json:"version"`
		Source    map[string]string `json:"source,omitempty"`
		Abandoned bool              `json:"abandoned"`
	}

	type lock struct {
		Packages []*pkg `json:"packages"`
	}

	lockObject := lock{
		Packages: make([]*pkg, count),
	}

	for i := range count {
		lockObject.Packages[i] = &pkg{
			Name:      fmt.Sprintf("vendor/package-%d", i),
			Version:   fmt.Sprintf("%d.%d.0", i%10, i%5),
			Source:    map[string]string{"reference": fmt.Sprintf("abc%d", i)},
			Abandoned: false,
		}
	}

	data, _ := json.Marshal(lockObject) //nolint:errchkjson // Not the purpose of the test

	return data
}
