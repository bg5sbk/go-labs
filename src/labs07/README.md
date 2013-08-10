测试几种链表遍历查找的性能差异。

实验结果：

    dada-imac:labs dada$ go test -test.bench="." labs07
    PASS
    Benchmark_Loop1    500000         4166 ns/op
    Benchmark_Loop2    500000         4983 ns/op
    Benchmark_Loop3    100000        14722 ns/op
    Benchmark_Loop4    500000         5068 ns/op
    ok  	labs07	8.891s

确定最后一种方式不会导致外部修改数据：

    dada-imac:labs dada$ go test -v labs07
    === RUN Test_Loop4
    --- PASS: Test_Loop4 (0.00 seconds)
    PASS
    ok  	labs07	0.014s

结论：既要防止外部回调函数不会误改数据，又要有好的遍历性能，试用最后一种方式。