package writing_benchmarks

import (
	"testing"
	"time"
)

func BenchmarkNoRepeat(b *testing.B) {
	// incorrect way to write a benchmark
	for i := 0; i < 1000000; i++ {
		time.Now()
	}
}

func add(a, b int) int { return a*a + b*b }

func BenchmarkOptimizedAway(b *testing.B) {
	// incorrect way to write a benchmark
	for i := 0; i < b.N; i++ {
		add(5, 7)
	}
}
