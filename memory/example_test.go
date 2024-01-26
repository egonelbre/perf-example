package example

import (
	"math/rand"
	"testing"
)

// More details at https://igoro.com/archive/gallery-of-processor-cache-effects/

var (
	data128B = [128 / 4]int32{}              // 128 bytes
	data1MB  = [1024 * 1024 / 4]int32{}      // 1 MB
	data64MB = [64 * 1024 * 1024 / 4]int32{} // 64MB

	order []int
)

func init() {
	order = make([]int, 1e6)
	for i := range order {
		order[i] = rand.Int()
	}
}

var sink int32

func Benchmark128B(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var total int32
		for _, k := range order {
			total += data128B[k%len(data128B)]
		}
		sink = total
	}
}

func Benchmark1MB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var total int32
		for _, k := range order {
			total += data1MB[k%len(data1MB)]
		}
		sink = total
	}
}

func Benchmark64MB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var total int32
		for _, k := range order {
			total += data64MB[k%len(data64MB)]
		}
		sink = total
	}
}
