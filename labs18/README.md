测试switch和回调函数效率差异。

实验结果：

```
dada-imac:labs18 dada$ go test -bench=".*"
testing: warning: no tests to run
PASS
Benchmark_UseSwitch	200000000	         9.58 ns/op
Benchmark_UseCallback	200000000	         8.05 ns/op
ok  	github.com/idada/go-labs/labs18	5.323s
```
