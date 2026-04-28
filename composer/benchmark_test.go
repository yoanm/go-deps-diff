package composer_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/yoanm/go-deps-diff/composer"
)

func BenchmarkLock_100Packages(b *testing.B) {
	data := generateLockFile(100)

	b.ResetTimer()

	for range b.N {
		if _, err := composer.ParseLock(data); err != nil {
			b.Fatalf("ParseLock failed: %v", err)
		}
	}
}

func BenchmarkLock_1000Packages(b *testing.B) {
	data := generateLockFile(1000)

	b.ResetTimer()

	for range b.N {
		if _, err := composer.ParseLock(data); err != nil {
			b.Fatalf("ParseLock failed: %v", err)
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

func BenchmarkReq_100Packages(b *testing.B) {
	data := generateReqFile(100)

	b.ResetTimer()

	for range b.N {
		if _, err := composer.ParseReq(data); err != nil {
			b.Fatalf("ParseReq failed: %v", err)
		}
	}
}

func BenchmarkReq_1000Packages(b *testing.B) {
	data := generateReqFile(1000)

	b.ResetTimer()

	for range b.N {
		if _, err := composer.ParseReq(data); err != nil {
			b.Fatalf("ParseReq failed: %v", err)
		}
	}
}

func generateReqFile(count int) []byte {
	req := make(map[string]string, count)

	for i := range count {
		req[fmt.Sprintf("vendor/package-%d", i)] = fmt.Sprintf("^%d.%d", i%10, i%5)
	}

	data, _ := json.Marshal(req) //nolint:errchkjson // Not the purpose of the test

	return data
}
