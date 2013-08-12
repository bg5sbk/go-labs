测试匿名函数和普通函数的调用消耗。

禁用优化前：

    dada-imac:labs dada$ go test -test.bench="." labs09
    testing: warning: no tests to run
    PASS
    Benchmark_NormalFuncCall    2000000000         0.52 ns/op
    Benchmark_VarFuncCall       2000000000         1.82 ns/op
    ok  	labs09	4.940s

禁用优化后：

    dada-imac:labs dada$ go test -test.bench="." -gcflags '-N' labs09
    testing: warning: no tests to run
    PASS
    Benchmark_NormalFuncCall    2000000000         1.98 ns/op
    Benchmark_VarFuncCall       1000000000         2.09 ns/op
    ok  	labs09	6.492s

结论：测试中用的是一个很简单的函数，所以很容易被新版的go做内联优化，禁用优化前后结果差别明显，如果不考虑优化，匿名函数和普通函数调用是一个量级的消耗，差别甚微。但是普通函数比较有可能被优化。