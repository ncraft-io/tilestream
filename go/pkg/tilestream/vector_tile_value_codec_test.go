package tilestream

import "testing"

func BenchmarkAppend(b *testing.B) {
	var a = []byte{1, 2, 3, 4, 5, 6, 7, 8}
	for i := 0; i < 100000000; i++ {
		appendVarint(a, 87654321)
	}
}
