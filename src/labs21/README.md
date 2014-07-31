Go的runtime包中使用goc机制实现Go的底层结构，所以判断goc机制应该不需要像cgo调用那样做防阻塞的保护，调用效率应该会高很多，这里做了一个测试goc机制调用c的实验。

实验结果：
$ go test labs21 --bench="."
testing: warning: no tests to run
PASS
Benchmark_Go_Call_C1	1000000000	         2.25 ns/op
ok  	labs21	2.489s

对比于labs20中的测试结果，效率的确高了很多，但是goc机制不允许包含标准c的头文件，是一个比较麻烦的事情，但是假设要用go实现一个脚本语言或调用脚本语言，用goc机制会获得更高的执行效率，当然实现起来代价也高了很多就是了。

如果可以找到一个快速把C的调用转换成类似runtime库中的.s文件那样的汇编代码，就可以做到无代价的包装C的库。