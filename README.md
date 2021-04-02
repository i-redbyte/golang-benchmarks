# Benchmark all #
## Benchmarking in my favorite Golang ##

Here I will add various benchmarks that I encounter while programming in Go.

**Example run all benchmarks:**
```
> go test -bench=. -benchmem -benchtime=256x
```

**Sample response:**
```
goos: darwin
goarch: amd64
pkg: benching
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
BenchmarkInsertIntMap100000-16               256             32162 ns/op           22539 B/op         15 allocs/op
BenchmarkInsertIntMap10000-16                256              3074 ns/op            2693 B/op          1 allocs/op
BenchmarkInsertIntMap1000-16                 256               223.3 ns/op           332 B/op          0 allocs/op
BenchmarkInsertIntMap100-16                  256                22.50 ns/op           20 B/op          0 allocs/op
```

**Example run selected benchmark:**
```
go test -bench=BenchmarkInsertIntMap10000 -benchmem -benchtime=256x
```
**Flags:**
- -benchmem - show memory usage
- -benchtime= time or number of times if x is at the end

[Sea more](https://golang.org/pkg/testing/)