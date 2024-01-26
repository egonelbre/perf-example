package example

// Documentation:
//
//    https://go.dev/doc/asm
//

// To build a single package:
//
//    go build -o example.o .

// To view assembly from compilation:
//
//    go tool compile -S example.go

// go tool objdump -S -s Add example.o
func Add(a, b int) int {
	return a + b
}

// go tool objdump -S -s Loop example.o
func Loop(a, b, n int) (r int) {
	for i := 0; i < n; i++ {
		r += a + b
	}
	return r
}

// Also:
//
//    go install loov.dev/lensm@main
//    lensm example.o
//
// If you want to properly understand assembly,
// then go write a virtual machine first.
//
// https://adventofcode.com/2019 is a pretty good thing to work through.
