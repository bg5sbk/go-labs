测试不修改runtime代码的情况下获取goid的效率。

测试结果:

```
$ go test -bench="."
testing: warning: no tests to run
PASS
Benchmark_Goid	  500000	      3612 ns/op
ok  	github.com/idada/go-labs/labs27	1.848s
```