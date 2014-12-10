测试平方根算法的效率。

测试结果：

```
dada-imac:labs dada$ go test labs16 --bench="."
testing: warning: no tests to run
PASS
Benchmark_Method1         500000              6677 ns/op
Benchmark_Method2       500000000             6.48 ns/op
Benchmark_Method3       500000000             3.73 ns/op
ok      labs16  9.558s
```