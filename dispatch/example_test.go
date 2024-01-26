package example

import (
	"math"
	"math/rand"
	"slices"
	"sort"
	"testing"
	"unsafe"
)

var unsorted []Shape
var sorted []Shape
var nointerface []ShapeStruct

func init() {
	unsorted = make([]Shape, 1e4)
	nointerface = make([]ShapeStruct, len(unsorted))

	for i := range unsorted {
		if rand.Intn(2) == 0 {
			unsorted[i] = Circle{rand.Float32()}
			nointerface[i] = ShapeStruct{CircleKind, rand.Float32()}
		} else {
			unsorted[i] = Square{rand.Float32()}
			nointerface[i] = ShapeStruct{SquareKind, rand.Float32()}
		}
	}

	sorted = slices.Clone(unsorted)
	type iface struct {
		itab uintptr
		data unsafe.Pointer
	}
	sort.Slice(sorted, func(i, k int) bool {
		a := (*iface)(unsafe.Pointer(&sorted[i])).itab
		b := (*iface)(unsafe.Pointer(&sorted[k])).itab
		return a < b
	})
}

type Shape interface {
	Area() float32
}

type Circle struct{ Radius float32 }
type Square struct{ Side float32 }

func (s Circle) Area() float32 {
	return math.Pi * s.Radius * s.Radius
}
func (s Square) Area() float32 {
	return s.Side * s.Side
}

func TotalArea(shapes []Shape) (total float32) {
	for _, shape := range shapes {
		total += shape.Area()
	}
	return total
}

func BenchmarkUnsorted(b *testing.B) {
	total := float32(0)
	for k := 0; k < b.N; k++ {
		total += TotalArea(unsorted)
	}
}

func BenchmarkSorted(b *testing.B) {
	total := float32(0)
	for k := 0; k < b.N; k++ {
		total += TotalArea(sorted)
	}
}

type ShapeKind byte

const (
	CircleKind = ShapeKind(0)
	SquareKind = ShapeKind(1)
)

type ShapeStruct struct {
	Kind ShapeKind
	Dim  float32
}

func (s ShapeStruct) Area() float32 {
	switch s.Kind {
	case CircleKind:
		return math.Pi * s.Dim * s.Dim
	case SquareKind:
		return s.Dim * s.Dim
	}
	return 0
}

func BenchmarkNoInterface(b *testing.B) {
	total := float32(0)
	for k := 0; k < b.N; k++ {
		for _, shape := range nointerface {
			total += shape.Area()
		}
	}
}
