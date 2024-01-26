package example

import (
	"testing"
)

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
func DotUnroll(xs []float32, incx int, ys []float32, incy int, n int) float32 {
	var r float32
	xi, yi := 0, 0
	for ; n >= 4; n -= 4 {
		r += xs[xi] * ys[yi]
		xi += incx
		yi += incy

		r += xs[xi] * ys[yi]
		xi += incx
		yi += incy

		r += xs[xi] * ys[yi]
		xi += incx
		yi += incy

		r += xs[xi] * ys[yi]
		xi += incx
		yi += incy
	}
	for ; n > 0; n-- {
		r += xs[xi] * ys[yi]

		xi += incx
		yi += incy
	}
	return r
}

//go:noinline
func DotPipeline(xs []float32, incx int, ys []float32, incy int, n int) float32 {
	var r1, r2, r3, r4 float32
	xi, yi := 0, 0
	for ; n >= 4; n -= 4 {
		r1 += xs[xi] * ys[yi]
		xi += incx
		yi += incy

		r2 += xs[xi] * ys[yi]
		xi += incx
		yi += incy

		r3 += xs[xi] * ys[yi]
		xi += incx
		yi += incy

		r4 += xs[xi] * ys[yi]
		xi += incx
		yi += incy
	}
	for ; n > 0; n-- {
		r1 += xs[xi] * ys[yi]

		xi += incx
		yi += incy
	}
	return r1 + r2 + r3 + r4
}

var sink float32

func BenchmarkDot(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sink += Dot(xs, incx, ys, incy, len(xs))
	}
}

func BenchmarkDotUnroll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sink += DotUnroll(xs, incx, ys, incy, len(xs))
	}
}

func BenchmarkDotPipeline(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sink += DotPipeline(xs, incx, ys, incy, len(xs))
	}
}
