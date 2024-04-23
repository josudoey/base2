# Base2

## Benchmark

```sh
$ go test -v -bench=. -run=none .
goos: darwin
goarch: arm64
pkg: github.com/josudoey/base2
BenchmarkFmtToString
BenchmarkFmtToString-8            863184              1358 ns/op
BenchmarkEncodeToString
BenchmarkEncodeToString-8        8685139               136.8 ns/op
PASS
ok      github.com/josudoey/base2       3.341s
```