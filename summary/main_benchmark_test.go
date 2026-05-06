package summary_test

import (
	"log/slog"
	"testing"

	"github.com/yoanm/go-deps-diff/summary"
)

func BenchmarkGenerateForChanges(b *testing.B) {
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() { // pb.Next() returns false when the benchmark should stop
			// Enable debug logs to check speed and allocation even at this level
			slog.SetLogLoggerLevel(slog.LevelDebug)

			_ = summary.GenerateForChanges(_integrationFullChanges)
		}
	})
}
