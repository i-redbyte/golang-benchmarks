package go_benchmark

import "testing"

func insertXIntSlice(x int, b *testing.B) {
	testSlice := make([]int, 0)
	b.ResetTimer()
	for i := 0; i < x; i++ {
		testSlice = append(testSlice, i)
	}
}

func BenchmarkInsertIntSlice1000000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXIntSlice(1000000, b)
	}
}

func BenchmarkInsertIntSlice100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXIntSlice(100000, b)
	}
}

func BenchmarkInsertIntSlice10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXIntSlice(10000, b)
	}
}

func BenchmarkInsertIntSlice1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXIntSlice(1000, b)
	}
}

func BenchmarkInsertIntSlice100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXIntSlice(100, b)
	}
}
