测试range循环和for循环，以及结构体循环和指针循环的性能区别。

    dada-imac:misc dada$ go test -test.bench="." labs04
    testing: warning: no tests to run
    PASS
    Benchmark_Loop1     2000000         923 ns/op
    Benchmark_Loop2     2000000         819 ns/op
    Benchmark_Loop3     2000000         825 ns/op
    Benchmark_Loop4     100000        26230 ns/op
    ok  	labs04	10.640s

结论：对结构体列表的range循环最消耗性能，因为数据要重复复制。