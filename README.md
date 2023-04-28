## Clean Code Performance Benchmark

Implement in Go the bechmarks presented in: 
	- https://www.youtube.com/watch?v=tD5NrevFtbU&t=19
	- https://www.computerenhance.com/p/clean-code-horrible-performance

### Results

```
goos: linux
goarch: amd64
pkg: clean-code
cpu: Intel(R) Core(TM) i7-8750H CPU @ 2.20GHz
BenchmarkAreaNewInterface1-12           143669847                8.434 ns/op
BenchmarkAreaEnum1-12                   271255455                4.398 ns/op
BenchmarkAreaNewInterface1000-12          143413              7452 ns/op
BenchmarkAreaEnum1000-12                  356358              3263 ns/op
PASS
ok      clean-code      6.055s
```

