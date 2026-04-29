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

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() { // pb.Next() returns false when the benchmark should stop
			if _, err := composer.ParseLock(data); err != nil {
				b.Fatalf("ParseLock failed: %v", err)
			}
		}
	})
}

func BenchmarkLock_1000Packages(b *testing.B) {
	data := generateLockFile(1000)

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() { // pb.Next() returns false when the benchmark should stop
			if _, err := composer.ParseLock(data); err != nil {
				b.Fatalf("ParseLock failed: %v", err)
			}
		}
	})
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

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() { // pb.Next() returns false when the benchmark should stop
			if _, err := composer.ParseReq(data); err != nil {
				b.Fatalf("ParseReq failed: %v", err)
			}
		}
	})
}

func BenchmarkReq_1000Packages(b *testing.B) {
	data := generateReqFile(1000)

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() { // pb.Next() returns false when the benchmark should stop
			if _, err := composer.ParseReq(data); err != nil {
				b.Fatalf("ParseReq failed: %v", err)
			}
		}
	})
}

func generateReqFile(count int) []byte {
	type reqStruct struct {
		Require    map[string]string `json:"require"`
		RequireDev map[string]string `json:"require-dev"`
	}

	reqObject := reqStruct{
		Require:    map[string]string{},
		RequireDev: map[string]string{},
	}

	for i := range count {
		if i%2 == 0 {
			reqObject.Require[fmt.Sprintf("vendor/package-%d", i)] = fmt.Sprintf("^%d.%d", i%10, i%5)
		} else {
			reqObject.RequireDev[fmt.Sprintf("vendor/package-%d", i)] = fmt.Sprintf("^%d.%d", i%10, i%5)
		}
	}

	data, _ := json.Marshal(reqObject) //nolint:errchkjson // Not the purpose of the test

	return data
}
