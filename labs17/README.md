尝试优化一段代码。

测试结果：

```
dada-imac:labs dada$ go test labs17 --bench="."
PASS
Benchmark_Method1          50000             53318 ns/op
Benchmark_Method2         200000             11122 ns/op
Benchmark_Method3         200000              9298 ns/op
Benchmark_Method4         200000              8838 ns/op
Benchmark_Method5         200000              8721 ns/op
ok      labs17  11.189s
```

还可以再继续优化初始时候的数值转字符串，数值转字符串也是比较大的消耗。