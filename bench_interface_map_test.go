package go_benchmark

import "testing"

func insertXInterfaceMap(x int, b *testing.B) {
	testmap := make(map[interface{}]int, 0)
	b.ResetTimer()
	for i := 0; i < x; i++ {
		testmap[i] = i
	}
}
func BenchmarkInsertInterfaceMap1000000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXInterfaceMap(1000000, b)
	}
}

func BenchmarkInsertInterfaceMap100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXInterfaceMap(100000, b)
	}
}

func BenchmarkInsertInterfaceMap10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXInterfaceMap(10000, b)
	}
}

func BenchmarkInsertInterfaceMap1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXInterfaceMap(1000, b)
	}
}

func BenchmarkInsertInterfaceMap100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXInterfaceMap(100, b)
	}
}
