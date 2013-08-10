同labs07，区别是用slice替代链表结构。

试验结果：

    dada-imac:labs dada$ go test -test.bench="." labs08
    PASS
    Benchmark_Loop1    2000000           837 ns/op
    Benchmark_Loop2    1000000          2671 ns/op
    Benchmark_Loop3     100000         15072 ns/op
    Benchmark_Loop4      10000        130577 ns/op
    Benchmark_Loop5       5000        606500 ns/op
    Benchmark_Loop6     500000          7145 ns/op
    ok  	labs08	    4.959s
    dada-imac:labs dada$ go test -test.bench="." labs07
    PASS
    Benchmark_Loop1    500000          4386 ns/op
    Benchmark_Loop2    500000          5206 ns/op
    Benchmark_Loop3    100000         14821 ns/op
    Benchmark_Loop4     10000        135384 ns/op
    Benchmark_Loop5      5000        607735 ns/op
    Benchmark_Loop6    500000          7067 ns/op
    ok  	labs07	14.618s    

结论：简单循环slice明显更快，但是加入查询后， slice的优势就可以忽略不计了。
