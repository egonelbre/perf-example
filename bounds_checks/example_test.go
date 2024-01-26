package example

import (
	"testing"
	"unsafe"
)

// To disable bounds checks entirely:
//
//    go test -gcflags=-B -bench .
//
// Viewing bounds checks
//
//    go test -gcflags "all=-m -m -d=ssa/check_bce/debug" -bench . 2>analysis.log
//    go install github.com/loov/view-annotated-file@latest
//    view-annotated-file analysis.log

var (
	xs   = make([]float32, 10000)
	ys   = make([]float32, 10000)
	incx = 1
	incy = 1
)

//go:noinline
func Dot(xs []float32, incx int, ys []float32, incy int, n int) float32 {
	var r float32
	xi, yi := 0, 0
	for ; n > 0; n-- {
		r += xs[xi] * ys[yi]

		xi += incx
		yi += incy
	}
	return r
}

//go:noinline
func DotUnsafe(xs []float32, incx int, ys []float32, incy int, n int) float32 {
	var r float32
	xi, yi := 0, 0
	for ; n > 0; n-- {
		r += *unsafeAt(xs, xi) * *unsafeAt(ys, yi)

		xi += incx
		yi += incy
	}
	return r
}

//go:noinline
func DotPointers(xs []float32, incx int, ys []float32, incy int, n int) float32 {
	var r float32
	xp := unsafe.Pointer(unsafe.SliceData(xs))
	yp := unsafe.Pointer(unsafe.SliceData(ys))
	incxp, incyp := uintptr(incx*4), uintptr(incy*4)
	for ; n > 0; n-- {
		r += *(*float32)(xp) * *(*float32)(yp)
		xp = unsafe.Add(xp, incxp)
		yp = unsafe.Add(yp, incyp)
	}
	return r
}

var sink float32

func BenchmarkDot(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sink += Dot(xs, incx, ys, incy, len(xs))
	}
}

func BenchmarkDotUnsafe(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sink += DotUnsafe(xs, incx, ys, incy, len(xs))
	}
}

func BenchmarkDotPointers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sink += DotPointers(xs, incx, ys, incy, len(xs))
	}
}

func unsafeAt[T any](xs []T, index int) *T {
	return (*T)(unsafe.Add(unsafe.Pointer(unsafe.SliceData(xs)), uintptr(index)*unsafe.Sizeof(xs[0])))
}
