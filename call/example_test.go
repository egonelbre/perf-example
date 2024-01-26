package example

import (
	"math/rand"
	"testing"
)

var sink int

func Nop() int {
	return 0
}

func BenchmarkInlined(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sink += Nop()
	}
}

//go:noinline
func Nop2() int {
	return 0
}

func BenchmarkNotInlined(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sink += Nop2()
	}
}

//go:noinline
func createANop() func() int { return Nop2 }

func BenchmarkFuncCall(b *testing.B) {
	nop := createANop()
	for i := 0; i < b.N; i++ {
		sink += nop()
	}
}

type Noper interface {
	Nop() int
}

type nop struct{}

func (nop) Nop() int { return 0 }

type nop2 struct{}

func (nop2) Nop() int { return 1 }

func BenchmarkDevirtualizedInterfaceCall(b *testing.B) {
	var nop Noper = nop{}
	for i := 0; i < b.N; i++ {
		sink += nop.Nop()
	}
}

func oneof() Noper {
	if rand.Intn(2) == 0 {
		return nop{}
	} else {
		return nop2{}
	}
}

func BenchmarkInterfaceCall(b *testing.B) {
	var nop Noper = oneof()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sink += nop.Nop()
	}
}

func BenchmarkInterfaceCall2(b *testing.B) {
	var nop Noper = oneof()
	b.ResetTimer()
	fn := nop.Nop
	for i := 0; i < b.N; i++ {
		sink += fn()
	}
}

func BenchmarkC(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sink += CNop()
	}
}
