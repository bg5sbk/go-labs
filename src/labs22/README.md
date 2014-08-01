尝试把CGO生成的runtime·cgocall替换为runtime·asmcgocall，调用开销显著下降。

实验过程和结果：

```
$ cd mylib/
$ go install
$ cd ..
$ go test -bench="."
PASS
Benchmark_Go_Call_C	50000000	        74.1 ns/op
ok  	labs23	3.782s
$ go test -bench="."
PASS
Benchmark_Go_Call_C	50000000	        54.7 ns/op
ok  	labs23	2.805s
$ cd mylib/
$ ./hack.sh
$ cd ..
$ go test -bench="."
PASS
Benchmark_Go_Call_C	100000000	        11.8 ns/op
ok  	labs23	1.203s
$ go test -bench="."
PASS
Benchmark_Go_Call_C	100000000	        12.2 ns/op
ok  	labs23	1.247s
```

