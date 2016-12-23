做内存池的时候有这样一个需求，给定一个数值，需要计算出这个数值落在2的第N次幂区间内，从而用这个N值索引出对应的池。

最直接的办法就是循环计算，因为2的20次幂就达到1M的大小了，这么大的内存块其实没有缓存的必要了，所以循环最多也就20次。

但是我很蛋疼了测试了sort包的Search函数，switch循环展开、if循环展开、人肉折半查找、浮点数取指数。

我电脑上测试结果如下：


go version 1.7.3

$ go test -bench="."
PASS`
Benchmark_Normal-8    	100000000	        22.8 ns/op
Benchmark_Search-8    	100000000	        21.5 ns/op
Benchmark_Search2-8   	200000000	         6.62 ns/op
Benchmark_Switch-8    	200000000	         7.54 ns/op
Benchmark_IF1-8       	300000000	         4.33 ns/op
Benchmark_IF2-8       	500000000	         2.71 ns/op
Benchmark_IF3-8       	2000000000	         1.09 ns/op
Benchmark_IF4-8       	200000000	         7.12 ns/op
Benchmark_IF5-8       	2000000000	         0.61 ns/op
PASS
ok  	command-line-arguments	18.027s


呵呵，蛋好疼。

补充：最后一个IF3测试是由群里的v1zze同学提供的算法，果然牛逼。
IF4 和 IF5 使用浮点数取指数计
