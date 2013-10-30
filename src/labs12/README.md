测试jsmalloc和malloc在go程序中是否有性能差别。

实验结果：

    dada-imac:labs dada$ go test -test.bench=".*" labs12
    PASS
    Benchmark_AllocAndFree	10000000	       190 ns/op
    Benchmark_JeAllocAndFree	20000000	       135 ns/op
    ok  	labs12	4.971s

实验结论：还是有差的