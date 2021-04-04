package go_benchmark

import "testing"

func selectXIntSlice(x int, b *testing.B) {
	testSlice := make([]int, x)
	for i := 0; i < x; i++ {
		testSlice[i] = i
	}
	var holder int
	b.ResetTimer()
	for i := 0; i < x; i++ {
		holder = testSlice[i]
	}
	if holder != 0 {

	}
}

func BenchmarkSelectIntSlice1000000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		selectXIntSlice(1000000, b)
	}
}

func BenchmarkSelectIntSlice100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		selectXIntSlice(100000, b)
	}
}

func BenchmarkSelectIntSlice10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		selectXIntSlice(10000, b)
	}
}

func BenchmarkSelectIntSlice1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		selectXIntSlice(1000, b)
	}
}

func BenchmarkSelectIntSlice100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		selectXIntSlice(100, b)
	}
}
