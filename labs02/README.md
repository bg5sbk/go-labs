测试值传参和指针传参的效率。

实验结果：

    dada-imac:misc dada$ go test -test.bench="." labs02
    testing: warning: no tests to run
    PASS
    Benchmark_Invoke1    2000000000          0.52 ns/op
    Benchmark_Invoke2    100000000           12.8 ns/op
    ok  	labs02	2.403s

结论：指针传参比值传参效率高很多。