package diff

import (
	"encoding/json"
	"fmt"
	"testing"
)

func BenchmarkDiff_100Packages(b *testing.B) {
	lockPrevious := generateLockFile(100)
	lockCurrent := generateLockFile(100)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Diff(lockPrevious, lockCurrent, nil, nil)
	}
}

func BenchmarkDiff_1000Packages(b *testing.B) {
	lockPrevious := generateLockFile(1000)
	lockCurrent := generateLockFile(1000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Diff(lockPrevious, lockCurrent, nil, nil)
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

	l := lock{
		Packages: make([]*pkg, count),
	}

	for i := 0; i < count; i++ {
		l.Packages[i] = &pkg{
			Name:      fmt.Sprintf("vendor/package-%d", i),
			Version:   fmt.Sprintf("%d.%d.0", i%10, i%5),
			Source:    map[string]string{"reference": fmt.Sprintf("abc%d", i)},
			Abandoned: false,
		}
	}

	data, _ := json.Marshal(l)
	return data
}
