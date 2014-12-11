测试interface{}类型转换判断和布尔值判断效率差异。

测试结果：

```
dada-imac:labs23 dada$ go test -v -bench=".*"
testing: warning: no tests to run
PASS
Benchmark_UseInterface	100000000	        10.4 ns/op
Benchmark_UseBoolean	100000000	        13.4 ns/op
ok  	github.com/idada/go-labs/labs23	2.414s
```
