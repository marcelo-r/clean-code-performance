## Clean Code Performance Benchmark

Implement in Go the bechmarks presented in: 
	- https://www.youtube.com/watch?v=tD5NrevFtbU&t=19
	- https://www.computerenhance.com/p/clean-code-horrible-performance

### Results

```
$ go test -benchmem -bench=.
goos: linux
goarch: amd64
pkg: github.com/marcelo-r/clean-code-performance
cpu: Intel(R) Core(TM) i7-8750H CPU @ 2.20GHz
BenchmarkAreaPoly1-12           147135076                7.925 ns/op           0 B/op          0 allocs/op
BenchmarkAreaEnum1-12           250384155                4.559 ns/op           0 B/op          0 allocs/op
BenchmarkAreaTable1-12          280415721                4.083 ns/op           0 B/op          0 allocs/op
BenchmarkAreaPoly1000-12          150037              8086 ns/op               0 B/op          0 allocs/op
BenchmarkAreaEnum1000-12          316341              3796 ns/op               0 B/op          0 allocs/op
BenchmarkAreaTable1000-12         428775              2772 ns/op               0 B/op          0 allocs/op
PASS
ok      github.com/marcelo-r/clean-code-performance     8.968s
```

