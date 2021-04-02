package go_benchmark

import "testing"

func insertXPreallocIntSlice(x int, b *testing.B) {
	testSlice := make([]int, x)
	b.ResetTimer()
	for i := 0; i < x; i++ {
		testSlice[i] = i
	}
}

func BenchmarkInsertIntSlicePreAllooc1000000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXPreallocIntSlice(1000000, b)
	}
}

func BenchmarkInsertIntSlicePreAllooc100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXPreallocIntSlice(100000, b)
	}
}

func BenchmarkInsertIntSlicePreAllooc10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXPreallocIntSlice(10000, b)
	}
}

func BenchmarkInsertIntSlicePreAllooc1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXPreallocIntSlice(1000, b)
	}
}

func BenchmarkInsertIntSlicePreAllooc100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXPreallocIntSlice(100, b)
	}
}
