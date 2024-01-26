package profiling

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

// go test -bench Label -benchmem
//
// go test -bench Label -cpuprofile cpu.prof -benchtime 5s
// go test -bench Label -memprofile mem.prof -benchtime 5s
// go test -bench Label -memprofile mem.prof -benchtime 5s
//
// On Non-Windows:
//    go tool pprof cpu.prof
//    go tool pprof mem.prof
//
//    go tool pprof -lines cpu.prof
//
// On Windows:
//    go tool pprof profiling.test.exe cpu.prof
//    go tool pprof profiling.test.exe mem.prof
//
//    go tool pprof -lines profiling.test.exe cpu.prof
//
// Commands in pprof:
//   top 30
//   top 30 -cum
//   list Format
//   disasm Format
//

var sink string

//go:noinline
func Format(prefix string, count int, suffix string) string {
	return fmt.Sprintf("%v%v%v", prefix, count, suffix)
}

//go:noinline
func Add(prefix string, count int, suffix string) string {
	return prefix + strconv.Itoa(count) + suffix
}

//go:noinline
func Builder(prefix string, count int, suffix string) string {
	var b strings.Builder
	b.Grow(len(prefix) + 13 + len(suffix))
	b.WriteString(prefix)

	var buffer [13]byte
	result := strconv.AppendInt(buffer[:], int64(count), 10)
	b.Write(result)

	b.WriteString(suffix)

	return b.String()
}

func BenchmarkLabel(b *testing.B) {
	b.Run("Format", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sink = Format("Alpha", i, "Variant")
		}
	})

	b.Run("Add", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sink = Add("Alpha", i, "Variant")
		}
	})

	b.Run("Builder", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sink = Builder("Alpha", i, "Variant")
		}
	})
}
