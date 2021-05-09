```
❯ go test -bench=.                     
goos: darwin
goarch: amd64
pkg: gopl/ch2/ex3
cpu: Intel(R) Core(TM) i5-8259U CPU @ 2.30GHz
BenchmarkPopCount-8             1000000000               0.2928 ns/op
BenchmarkPopCountByLoop-8       266833170                4.499 ns/op
PASS
ok      gopl/ch2/ex3    2.513s
```