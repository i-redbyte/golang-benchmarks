package go_benchmark

import "testing"

/*
console run example:
go test -bench=BenchmarkInsertIntMap -benchmem -benchtime=256x
*/

func insertXPreallocIntMap(x int, b *testing.B) {
	testmap := make(map[int]int, x)
	b.ResetTimer()
	for i := 0; i < x; i++ {
		testmap[i] = i
	}
}

func BenchmarkInsertIntMapPreAlloc1000000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXPreallocIntMap(1000000, b)
	}
}

func BenchmarkInsertIntMapPreAlloc100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXPreallocIntMap(100000, b)
	}
}

func BenchmarkInsertIntMapPreAlloc10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXPreallocIntMap(10000, b)
	}
}

func BenchmarkInsertIntMapPreAlloc1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXPreallocIntMap(1000, b)
	}
}

func BenchmarkInsertIntMapPreAlloc100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXPreallocIntMap(100, b)
	}
}
