测试binary.Write和硬编码的效率差异。

测试结果：

```
dada-imac:labs24 dada$ go test -v -bench='.*'
testing: warning: no tests to run
PASS
Benchmark_UseBinaryWrite1	50000000	        68.5 ns/op
Benchmark_UseBinaryWrite2	10000000	       279 ns/op
Benchmark_UseHardcode	100000000	        21.5 ns/op
ok  	github.com/idada/go-labs/labs24	8.771s
```