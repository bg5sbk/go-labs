测试LockOSThread()对chan消息处理的影响，本来是希望通过LockOSThread()可以独占线程，避免chan通讯发生的调度开销，结果跟预期相反，由于调度算法没有改变，LockOSThread()的goroutine反而需要等待关联的线程空闲了才能被执行到，所以反而是变慢了。

测试结果：

```
$ go test -bench="."
testing: warning: no tests to run
PASS
Benchmark_Normal	10000000	       220 ns/op
Benchmark_LockThread	 1000000	      2084 ns/op
ok  	github.com/idada/go-labs/labs25	4.540s
```