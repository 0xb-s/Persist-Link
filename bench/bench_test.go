package benchmark

import (
	"persistlink/lib"
	"testing"
)

func BenchmarkAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		link := lib.Empty[int]()
		for j := 0; j < 1000; j++ {
			link = link.Append(j)
		}
	}
}

func BenchmarkConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		linkA := lib.Empty[int]()
		linkB := lib.Empty[int]()
		for j := 0; j < 500; j++ {
			linkA = linkA.Append(j)
		}
		for j := 500; j < 1000; j++ {
			linkB = linkB.Append(j)
		}
		_ = linkA.Concat(linkB)
	}
}

func BenchmarkAsSlice(b *testing.B) {
	link := lib.Empty[int]()
	for i := 0; i < 1000; i++ {
		link = link.Append(i)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = link.AsSlice()
	}
}
