package example

import (
	"math/rand"
	"slices"
	"testing"
)

var data []float32
var sorted []*float32
var unsorted []*float32

func init() {
	data = make([]float32, 10000000)
	for i := range data {
		data[i] = rand.Float32()
	}

	sorted = make([]*float32, len(data))
	for i := range sorted {
		sorted[i] = &data[i]
	}
	unsorted = slices.Clone(sorted)
	rand.Shuffle(len(unsorted), func(i, k int) {
		unsorted[i], unsorted[k] = unsorted[k], unsorted[i]
	})
}

var sink float32

func BenchmarkUnsorted(b *testing.B) {
	for k := 0; k < b.N; k++ {
		total := float32(0)
		for _, v := range unsorted {
			total += *v
		}
		sink += total
	}
}

func BenchmarkSorted(b *testing.B) {
	for k := 0; k < b.N; k++ {
		total := float32(0)
		for _, v := range sorted {
			total += *v
		}
		sink += total
	}
}

func BenchmarkData(b *testing.B) {
	for k := 0; k < b.N; k++ {
		total := float32(0)
		for _, v := range data {
			total += v
		}
		sink += total
	}
}
