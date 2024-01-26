package example

import (
	"testing"
)

type Shape struct {
	Kind      int
	Dimension float32

	ExtraData [10 * 1024]byte
}

var data = make([]Shape, 1024)

func BenchmarkCopy(b *testing.B) {
	total := float32(0)
	for k := 0; k < b.N; k++ {
		for i, shape := range data {
			total += shape.Dimension
			shape.Dimension++
			data[i] = shape
		}
	}
}

func BenchmarkReference(b *testing.B) {
	total := float32(0)
	for k := 0; k < b.N; k++ {
		for i := range data {
			shape := &data[i]
			total += shape.Dimension
			shape.Dimension++
		}
	}
}

func BenchmarkCall(b *testing.B) {
	total := float32(0)
	for k := 0; k < b.N; k++ {
		for _, shape := range data {
			total += DimensionValue(shape)
		}
	}
}

func BenchmarkCallPointer(b *testing.B) {
	total := float32(0)
	for k := 0; k < b.N; k++ {
		for i := range data {
			shape := &data[i]
			total += DimensionPointer(shape)
		}
	}
}

//go:noinline
func DimensionValue(v Shape) float32 { return v.Dimension }

//go:noinline
func DimensionPointer(v *Shape) float32 { return v.Dimension }
