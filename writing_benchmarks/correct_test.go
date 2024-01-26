package writing_benchmarks

import (
	"runtime"
	"testing"
	"time"

	"github.com/loov/hrtime"
)

var sinkTime time.Time
var sinkDuration time.Duration

func BenchmarkTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sinkTime = time.Now()
	}
}

func BenchmarkTimeAlternative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(time.Now())
	}
}

// Using https://pkg.go.dev/golang.org/x/perf/cmd/benchstat
//
//	go test -bench Sub -count 10 | tee bench.log
//	go install golang.org/x/perf/cmd/benchstat@latest
//	benchstat bench.log
//	benchstat old.log new.log
//	benchstat -col /impl -row .name bench.log
func BenchmarkSub(b *testing.B) {
	b.Run("impl=hrtime", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sinkDuration = hrtime.Now()
		}
	})

	b.Run("impl=time", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sinkTime = time.Now()
		}
	})
}
