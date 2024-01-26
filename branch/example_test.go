package example

import (
	"math/rand"
	"slices"
	"sort"
	"testing"
)

// More details at https://igoro.com/archive/fast-and-slow-if-statements-branch-prediction-in-modern-processors/

var unsorted []int
var sorted []int
var half []int

func init() {
	unsorted = make([]int, 10000)
	for i := range unsorted {
		unsorted[i] = rand.Intn(100)
	}

	sorted = slices.Clone(unsorted)
	sort.Ints(sorted)

	half = make([]int, len(unsorted))
	for i := range half {
		if i%2 == 0 {
			half[i] = 0
		} else {
			half[i] = 100
		}
	}
}

//go:noinline
func DiffLimit(vs []int, limit int) int {
	above := 0
	below := 0
	for _, v := range vs {
		if v > limit {
			above += v
		} else {
			below += v
		}
	}
	return above - below
}

func BenchmarkUnsorted(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DiffLimit(unsorted, 50)
	}
}

func BenchmarkSorted(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DiffLimit(sorted, 50)
	}
}

func BenchmarkHalf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DiffLimit(half, 50)
	}
}
