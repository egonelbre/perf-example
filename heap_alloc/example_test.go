package example

import (
	"testing"
	"encoding/binary"
)

var escape []byte

func Heap(v uint64) byte {
	escape = make([]byte, 8)
	binary.LittleEndian.PutUint64(escape[:], v)
	return escape[0]
}

func Stack(v uint64) byte {
	var data [8]byte
	binary.LittleEndian.PutUint64(data[:], v)
	return data[0]
}

var sink byte

func BenchmarkHeap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sink += Heap(uint64(i))
	}
}

func BenchmarkStack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sink += Stack(uint64(i))
	}
}
