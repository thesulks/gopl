```
❯ go test -bench=. -benchmem
goos: darwin
goarch: amd64
pkg: gopl/ch4/ex4
cpu: Intel(R) Core(TM) i5-8259U CPU @ 2.30GHz
BenchmarkRotateLeft-8                       1093           1091689 ns/op         4013210 B/op          1 allocs/op
BenchmarkRotateLeftByReverse-8              1053           1142125 ns/op            7600 B/op          0 allocs/op
PASS
ok      gopl/ch4/ex4    3.710s
```