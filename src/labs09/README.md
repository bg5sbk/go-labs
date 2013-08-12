测试匿名函数和普通函数的调用消耗。

禁用优化前：

    dada-imac:labs dada$ go test -test.bench="." labs09
    testing: warning: no tests to run
    PASS
    Benchmark_NormalFuncCall  2000000000         0.52 ns/op
    Benchmark_VarFuncCall1    2000000000         1.84 ns/op
    Benchmark_VarFuncCall2    1000000000         2.16 ns/op
    ok  	labs09	7.347s

禁用优化后：

    dada-imac:labs dada$ go test -test.bench="." -gcflags '-N' labs09
    testing: warning: no tests to run
    PASS
    Benchmark_NormalFuncCall  2000000000         1.99 ns/op
    Benchmark_VarFuncCall1    1000000000         2.02 ns/op
    Benchmark_VarFuncCall2    50000000           58.8 ns/op
    ok  	labs09	9.412s

结论：测试中用的是一个很简单的函数，所以很容易被新版的go做内联优化，禁用优化前后结果差别明显，如果不考虑优化，匿名函数和普通函数调用是一个量级的消耗，差别甚微。但是普通函数比较有可能被优化。在没有优化的情况下，闭包函数消耗又要再高一个量级。