package go_benchmark

import "testing"

func insertXIntElementMap(x int, b *testing.B) {
	testmap := make(map[int]int, 0)
	b.ResetTimer()
	for i := 0; i < x; i++ {
		testmap[i] = i
	}
}

func BenchmarkInsertIntMap100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXIntElementMap(100000, b)
	}
}

func BenchmarkInsertIntMap10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXIntElementMap(10000, b)
	}
}

func BenchmarkInsertIntMap1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXIntElementMap(1000, b)
	}
}

func BenchmarkInsertIntMap100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXIntElementMap(100, b)
	}
}
