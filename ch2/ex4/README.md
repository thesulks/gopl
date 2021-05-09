```
❯ go test -bench=.
goos: darwin
goarch: amd64
pkg: gopl/ch2/ex4
cpu: Intel(R) Core(TM) i5-8259U CPU @ 2.30GHz
BenchmarkPopCount-8             1000000000               0.2934 ns/op
BenchmarkBitCountByShift-8      40639585                27.65 ns/op
PASS
ok      gopl/ch2/ex4    1.989s
```