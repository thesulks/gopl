```console
❯ go test -bench=. -benchmem
goos: darwin
goarch: amd64
pkg: gopl/ch1/ex3
cpu: Intel(R) Core(TM) i5-8259U CPU @ 2.30GHz
BenchmarkEchoNaive10-8           2209930               559.8 ns/op           648 B/op          9 allocs/op
BenchmarkEchoNaive100-8            81615             14019 ns/op           58808 B/op         99 allocs/op
BenchmarkEchoNaive1000-8            1312            893401 ns/op         5827810 B/op        999 allocs/op
BenchmarkEchoNaive10000-8             16          63532415 ns/op        582209243 B/op     10022 allocs/op
BenchmarkEchoJoin10-8            9347233               125.0 ns/op           112 B/op          1 allocs/op
BenchmarkEchoJoin100-8            963727              1119 ns/op            1152 B/op          1 allocs/op
BenchmarkEchoJoin1000-8           104608             11154 ns/op           12288 B/op          1 allocs/op
BenchmarkEchoJoin10000-8            9253            110736 ns/op          114777 B/op          1 allocs/op
PASS
ok      gopl/ch1/ex3    10.610s
```