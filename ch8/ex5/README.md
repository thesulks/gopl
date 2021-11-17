```console
❯ go test -bench=.              
goos: darwin
goarch: amd64
pkg: gopl/ch8/ex5
cpu: Intel(R) Core(TM) i5-8259U CPU @ 2.30GHz
BenchmarkMandelbrotImage-8                                     5         203978827 ns/op
BenchmarkMandelbrotImageWaitGroup-8                            4         281439930 ns/op
BenchmarkMandelbrotImageWaitGroupWithToken10-8                 2         572686250 ns/op
BenchmarkMandelbrotImageWaitGroupWithToken100-8                3         476949541 ns/op
BenchmarkMandelbrotImageWaitGroupWithToken1000-8               3         467150924 ns/op
BenchmarkMandelbrotImageWaitGroupWithToken10000-8              3         471643164 ns/op
BenchmarkMandelbrotImageWaitGroupWithToken100000-8             3         487783074 ns/op
BenchmarkMandelbrotImageWaitGroupWithToken1000000-8            3         475682882 ns/op
BenchmarkMandelbrotImageLocality2-8                           10         106074651 ns/op
BenchmarkMandelbrotImageLocality4-8                           12          96196792 ns/op
BenchmarkMandelbrotImageLocality8-8                           13          78826974 ns/op
BenchmarkMandelbrotImageLocality16-8                          18          61761034 ns/op
BenchmarkMandelbrotImageLocality32-8                          21          58654179 ns/op
BenchmarkMandelbrotImageLocality64-8                          18          59655761 ns/op
BenchmarkMandelbrotImageLocality128-8                         19          56337867 ns/op
BenchmarkMandelbrotImageLocality256-8                         18          58008897 ns/op
BenchmarkMandelbrotImageLocality512-8                         20          55843906 ns/op
BenchmarkMandelbrotImageLocality1024-8                        18          90870795 ns/op
PASS
ok      gopl/ch8/ex5    36.675s
```

```
❯ go test -bench=.
goos: darwin
goarch: amd64
pkg: gopl/ch8/ex5
cpu: Intel(R) Core(TM) i5-8259U CPU @ 2.30GHz
BenchmarkMandelbrotImage-8                             5         204055164 ns/op
BenchmarkMandelbrotImageLocality2-8                   10         104483474 ns/op
BenchmarkMandelbrotImageLocality4-8                   12          94718173 ns/op
BenchmarkMandelbrotImageLocality8-8                   14          75438849 ns/op
BenchmarkMandelbrotImageLocality16-8                  19          57930228 ns/op
BenchmarkMandelbrotImageLocality32-8                  21          53371042 ns/op
BenchmarkMandelbrotImageLocality64-8                  22          51509369 ns/op
BenchmarkMandelbrotImageLocality128-8                 21          48411504 ns/op
BenchmarkMandelbrotImageLocality256-8                 22          49645759 ns/op
BenchmarkMandelbrotImageLocality512-8                 24          50816722 ns/op
BenchmarkMandelbrotImageLocality1024-8                18          57662734 ns/op
PASS
ok      gopl/ch8/ex5    14.998s
```

Is image.RGBA concurrency-safe?
