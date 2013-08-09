测试整数和浮点数的运算效率。

实验结果：

    dada-imac:labs dada$ go test -test.bench="." labs05
    testing: warning: no tests to run
    PASS
    Benchmark_IntAdd        2000000000           0.28 ns/op
    Benchmark_Int8Add       2000000000           0.28 ns/op
    Benchmark_Int16Add      2000000000           0.28 ns/op
    Benchmark_Int32Add      2000000000           0.28 ns/op
    Benchmark_Int64Add      2000000000           0.28 ns/op
    Benchmark_Float32Add    2000000000           1.05 ns/op
    Benchmark_Float64Add    2000000000           0.79 ns/op
    Benchmark_IntSub        2000000000           0.28 ns/op
    Benchmark_Int8Sub       2000000000           0.28 ns/op
    Benchmark_Int16Sub      2000000000           0.27 ns/op
    Benchmark_Int32Sub      2000000000           0.28 ns/op
    Benchmark_Int64Sub      2000000000           0.28 ns/op
    Benchmark_Float32Sub    2000000000           1.31 ns/op
    Benchmark_Float64Sub    2000000000           0.87 ns/op
    Benchmark_IntMul        2000000000           0.79 ns/op
    Benchmark_Int8Mul       2000000000           0.79 ns/op
    Benchmark_Int16Mul      2000000000           0.79 ns/op
    Benchmark_Int32Mul      2000000000           0.79 ns/op
    Benchmark_Int64Mul      2000000000           0.79 ns/op
    Benchmark_Float32Mul    2000000000           1.57 ns/op
    Benchmark_Float64Mul    2000000000           1.31 ns/op
    Benchmark_IntDiv        2000000000           1.57 ns/op
    Benchmark_Int8Div       2000000000           1.84 ns/op
    Benchmark_Int16Div      1000000000           2.08 ns/op
    Benchmark_Int32Div      2000000000           1.62 ns/op
    Benchmark_Int64Div      2000000000           1.57 ns/op
    Benchmark_Float32Div    50000000            44.3 ns/op
    Benchmark_Float64Div    50000000            48.0 ns/op
    ok  	labs05	49.453s    

结论：浮点数除法明显更慢，尽量把除法转换为乘法运算。