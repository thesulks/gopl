```
❯ go test -bench=.
goos: darwin
goarch: amd64
pkg: gopl/ch2/ex5
cpu: Intel(R) Core(TM) i5-8259U CPU @ 2.30GHz
BenchmarkPopCount-8                     1000000000               0.2940 ns/op
BenchmarkBitCountByClearing-8           94826704                12.76 ns/op
PASS
ok      gopl/ch2/ex5    1.651s
```