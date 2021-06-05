```
❯ go test -bench=. -benchmem
goos: darwin
goarch: amd64
pkg: gopl/ch4/ex1
cpu: Intel(R) Core(TM) i5-8259U CPU @ 2.30GHz
BenchmarkDiffBitCountSha256-8           43188356                27.37 ns/op            0 B/op          0 allocs/op
BenchmarkDiffBitCountSha256Ref-8        36294865                32.58 ns/op            0 B/op          0 allocs/op
PASS
ok      gopl/ch4/ex1    2.550s
```
