package go_benchmark

import "testing"

func selectXIntMap(x int, b *testing.B) {
	testmap := make(map[int]int, x)
	for i := 0; i < x; i++ {
		testmap[i] = i
	}
	var holder int
	b.ResetTimer()
	for i := 0; i < x; i++ {
		holder = testmap[i]
	}
	if holder != 0 {
		//little trick
	}
}

func BenchmarkSelectIntMap1000000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		selectXIntMap(1000000, b)
	}
}

func BenchmarkSelectIntMap100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		selectXIntMap(100000, b)
	}
}

func BenchmarkSelectIntMap10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		selectXIntMap(10000, b)
	}
}

func BenchmarkSelectIntMap1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		selectXIntMap(1000, b)
	}
}

func BenchmarkSelectIntMap100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		selectXIntMap(100, b)
	}
}
