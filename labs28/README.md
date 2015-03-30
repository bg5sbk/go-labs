测试`[]byte`转`string`的效率。

测试结果：

```
$ go test -bench="."
PASS
Benchmark_Normal	20000000	        63.4 ns/op
Benchmark_ByteString	2000000000	         0.55 ns/op
ok  	github.com/idada/go-labs/labs28	2.486s
```
