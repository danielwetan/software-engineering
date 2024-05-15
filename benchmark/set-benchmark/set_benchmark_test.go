package setbenchmark

import (
	"math/rand"
	"testing"
)

func benchmark(b *testing.B, set Set, size int, fill int) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < size; j++ {
			set.Add(rand.Int() % fill)
		}
	}
}

func BenchmarkSetWithFewCollisions(b *testing.B) {
	benchmark(b, NewSet(), 100, 700)
}

func BenchmarkRealtimeArraySetWithFewCollisions(b *testing.B) {
	benchmark(b, NewRealtimeArraySet(), 100, 700)
}

func BenchmarkOnDemandArraySetWithFewCollisions(b *testing.B) {
	benchmark(b, NewOnDemandArraySet(), 100, 700)
}
