测试Go调用C和C调用Go的效率差异。

测试结果：

> go test labs20 -bench=".*"
testing: warning: no tests to run
PASS
Benchmark_Go_Call_C	50000000	        55.6 ns/op
Benchmark_C_Call_GO	 5000000	       434 ns/op
ok  	labs20	5.473s
