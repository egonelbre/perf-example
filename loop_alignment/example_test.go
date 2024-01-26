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
func Dot1(xs []float32, incx int, ys []float32, incy int, n int) float32 {
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
func Dot2(xs []float32, incx int, ys []float32, incy int, n int) float32 {
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
func Dot3(xs []float32, incx int, ys []float32, incy int, n int) float32 {
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
func Dot4(xs []float32, incx int, ys []float32, incy int, n int) float32 {
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
func Dot5(xs []float32, incx int, ys []float32, incy int, n int) float32 {
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
func Dot6(xs []float32, incx int, ys []float32, incy int, n int) float32 {
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
func Dot7(xs []float32, incx int, ys []float32, incy int, n int) float32 {
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
func Dot8(xs []float32, incx int, ys []float32, incy int, n int) float32 {
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
func Dot9(xs []float32, incx int, ys []float32, incy int, n int) float32 {
	var r float32
	xi, yi := 0, 0
	for ; n > 0; n-- {
		r += xs[xi] * ys[yi]

		xi += incx
		yi += incy
	}
	return r
}

var sink float32

func BenchmarkDot1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sink += Dot1(xs, incx, ys, incy, len(xs))
	}
}

func BenchmarkDot2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sink += Dot2(xs, incx, ys, incy, len(xs))
	}
}

func BenchmarkDot3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sink += Dot3(xs, incx, ys, incy, len(xs))
	}
}

func BenchmarkDot4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sink += Dot4(xs, incx, ys, incy, len(xs))
	}
}

func BenchmarkDot5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sink += Dot5(xs, incx, ys, incy, len(xs))
	}
}

func BenchmarkDot6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sink += Dot6(xs, incx, ys, incy, len(xs))
	}
}

func BenchmarkDot7(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sink += Dot7(xs, incx, ys, incy, len(xs))
	}
}

func BenchmarkDot8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sink += Dot8(xs, incx, ys, incy, len(xs))
	}
}

func BenchmarkDot9(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sink += Dot9(xs, incx, ys, incy, len(xs))
	}
}
