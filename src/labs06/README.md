测试小数据量循环和map取数据以及硬编码取数据的性能消耗。

试验结果：

    dada-imac:labs dada$ go test -test.bench="." labs06
    testing: warning: no tests to run
    PASS
    Benchmark_Loop1    500000000             5.73 ns/op
    Benchmark_Loop2    500000000             5.72 ns/op
    Benchmark_Loop3    50000000              68.0 ns/op
    Benchmark_Loop4    500000000             4.92 ns/op
    Benchmark_Loop5    500000000             4.40 ns/op
    ok  	labs06	15.970s    

结论：硬编码 < 指针slice的range循环 < for循环，但是量级是一样的，看情况用。但是map差了一个量级，小数据量尽量少用。
