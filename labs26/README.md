比较直接调用函数和反射调用函数的效率差别，忘记为什么目的测试的了。

测试结果：

```
$ go test -bench="."
testing: warning: no tests to run
PASS
Benchmark_NormalFuncCall	2000000000	         0.53 ns/op
Benchmark_ReflectFuncCall	 2000000	       667 ns/op
ok  	github.com/idada/go-labs/labs26	3.116s
```