测试对象创建的效率。

    dada-imac:misc dada$ go test -test.bench="." labs03
    testing: warning: no tests to run
    PASS
    Benchmark_NewStruct1    100000000            13.0 ns/op
    Benchmark_NewStruct2    100000000            24.5 ns/op
    Benchmark_NewStruct3    50000000             37.8 ns/op
    Benchmark_NewStruct4    100000000            25.1 ns/op
    Benchmark_NewStruct5    200000000            8.65 ns/op
    ok  	labs03	10.872s

结论：5和4的结果比较意外。