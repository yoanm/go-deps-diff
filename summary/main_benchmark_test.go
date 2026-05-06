package summary_test

import (
	"io"
	"log"
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
			log.SetOutput(io.Discard) // But drop output to avoid parsing issues later

			_ = summary.GenerateForChanges(_integrationFullChanges)
		}
	})
}
