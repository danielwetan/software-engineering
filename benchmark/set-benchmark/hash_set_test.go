package setbenchmark

import (
	"testing"
)

func BenchmarkSmallSetWithFewCollisions(b *testing.B) {
	benchmark(b, NewSet(), 100, 700)
}
func BenchmarkSmallSetWithMoreCollisions(b *testing.B) {
	benchmark(b, NewSet(), 100, 100)
}
func BenchmarkSmallSetWithManyCollisions(b *testing.B) {
	benchmark(b, NewSet(), 100, 25)
}
func BenchmarkMediumSetWithFewCollisions(b *testing.B) {
	benchmark(b, NewSet(), 5000, 35000)
}
func BenchmarkMediumSetWithMoreCollisions(b *testing.B) {
	benchmark(b, NewSet(), 5000, 5000)
}
func BenchmarkMediumSetWithManyCollisions(b *testing.B) {
	benchmark(b, NewSet(), 5000, 1250)
}
func BenchmarkLargeSetWithFewCollisions(b *testing.B) {
	benchmark(b, NewSet(), 100000, 700000)
}
func BenchmarkLargeSetWithMoreCollisions(b *testing.B) {
	benchmark(b, NewSet(), 100000, 100000)
}
func BenchmarkLargeSetWithManyCollisions(b *testing.B) {
	benchmark(b, NewSet(), 100000, 25000)
}

// sample output for benchmark test
// goos: linux
// goarch: amd64
// pkg: setbenchmark
// cpu: AMD Ryzen 7 4700U with Radeon Graphics
// BenchmarkSmallSetWithFewCollisions-8     	   81745	     14085 ns/op	    3267 B/op	      15 allocs/op
// BenchmarkSmallSetWithMoreCollisions-8    	   90435	     12967 ns/op	    3119 B/op	      12 allocs/op
// BenchmarkSmallSetWithManyCollisions-8    	  151624	      7329 ns/op	     647 B/op	       5 allocs/op
// BenchmarkMediumSetWithFewCollisions-8    	    1660	    703493 ns/op	  213518 B/op	     186 allocs/op
// BenchmarkMediumSetWithMoreCollisions-8   	    2258	    503664 ns/op	  112361 B/op	     158 allocs/op
// BenchmarkMediumSetWithManyCollisions-8   	    2858	    404115 ns/op	   53525 B/op	      74 allocs/op
// BenchmarkLargeSetWithFewCollisions-8     	      76	  13822315 ns/op	 3525810 B/op	    3325 allocs/op
// BenchmarkLargeSetWithMoreCollisions-8    	      74	  15287934 ns/op	 3379528 B/op	    2335 allocs/op
// BenchmarkLargeSetWithManyCollisions-8    	     166	   6030087 ns/op	  889381 B/op	     921 allocs/op
// PASS
// coverage: 40.0% of statements
// ok  	setbenchmark	11.397s

// how to read
// BenchmarkSmallSetWithFewCollisions-8     	   81745	     14085 ns/op	    3267 B/op	      15 allocs/op

// Benchmark Name: BenchmarkSmallSetWithFewCollisions-8
// This is the name of the benchmark test. In this case, it seems to be testing something related to a small set with few collisions.
//
// Iterations: 81745
// This is the number of iterations performed during the benchmark test.
//
// Average Time per Operation: 14085 ns/op
// This indicates the average time taken per operation (in nanoseconds) during the benchmark. Lower is better, as it means the code is executing faster.
//
// Allocated Memory per Operation: 3267 B/op
// This is the amount of memory allocated per operation (in bytes). Lower is generally better, as it means the code is more memory-efficient.
//
// Allocations per Operation: 15 allocs/op
// This is the average number of allocations performed per operation. Fewer allocations are generally better for performance.

func TestCanAddUniqueValues(t *testing.T) {
	values := []int{5, 4, 3, 9, 1, 2, 8}
	set := NewSet()

	for _, v := range values {
		set.Add(v)
	}
	set.RemoveDuplicates()
	Assert(t, set.Length() == len(values), "Wrong length")

	for _, v := range values {
		Assert(t, set.Contains(v), "Missing value")
	}
}

func TestUniqueRemoveDuplicates(t *testing.T) {
	values := []int{5, 4, 3, 5, 2, 1, 2, 3, 4, 5, 1, 2, 2}
	uniques := []int{5, 4, 3, 2, 1}
	set := NewSet()

	for _, v := range values {
		set.Add(v)
	}
	set.RemoveDuplicates()

	Assert(t, set.Length() == len(uniques), "Wrong length")
	for _, v := range uniques {
		Assert(t, set.Contains(v), "Missing value")
	}
}

func Assert(t *testing.T, condition bool, message string) {
	if !condition {
		t.Fatal(message)
	}
}
