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
BenchmarkAreaInterface1-12              148383061                7.781 ns/op           0 B/op          0 allocs/op
BenchmarkAreaEnum1-12                   274662770                4.396 ns/op           0 B/op          0 allocs/op
BenchmarkAreaTable1-12                  285506127                4.018 ns/op           0 B/op          0 allocs/op
BenchmarkAreaInterface1000-12             159108              7812 ns/op               0 B/op          0 allocs/op
BenchmarkAreaEnum1000-12                  305194              3941 ns/op               0 B/op          0 allocs/op
BenchmarkAreaTable1000-12                 406488              2545 ns/op               0 B/op          0 allocs/op
PASS
ok      github.com/marcelo-r/clean-code-performance     9.791s
```

