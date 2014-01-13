测试缓存反射信息对效率的影响。

```
dada-imac:src dada$ go test labs19 --bench=".*" -v
=== RUN Test_Verify
--- PASS: Test_Verify (0.00 seconds)
=== RUN Test_FastVerify
--- PASS: Test_FastVerify (0.00 seconds)
=== RUN Test_VeryFastVerify
--- PASS: Test_VeryFastVerify (0.00 seconds)
PASS
Benchmark_________Verify	 1000000	      1399 ns/op
Benchmark_____FastVerify	10000000	       178 ns/op
Benchmark_VeryFastVerify	20000000	        93.9 ns/op
ok  	labs19	5.364s
```