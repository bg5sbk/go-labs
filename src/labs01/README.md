测试类型判断和类型转换的效率。

执行结果：

    dada-imac:misc dada$ go test -test.bench=".*" labs01
    testing: warning: no tests to run
    PASS
    Benchmark_TypeSwitch	50000000	        33.0 ns/op
    Benchmark_NormalSwitch	2000000000	         1.99 ns/op
    Benchmark_InterfaceSwitch	100000000	        18.4 ns/op
    ok  	labs	7.741s

结论：类型判断和类型转换这两个操作都比直接操作多几倍的消耗。

