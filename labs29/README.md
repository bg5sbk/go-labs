测试不同压缩算法压缩json数据的压缩比和压缩效率。

实验结果：

```
dada-imac:labs29 dada$ go test -v
=== RUN Test_Normal
--- PASS: Test_Normal (0.16s)
=== RUN Test_Gzip_Level1
--- PASS: Test_Gzip_Level1 (0.14s)
=== RUN Test_Gzip_Level5
--- PASS: Test_Gzip_Level5 (0.16s)
=== RUN Test_Gzip_Level9
--- PASS: Test_Gzip_Level9 (0.16s)
=== RUN Test_Snappy
--- PASS: Test_Snappy (0.18s)
=== RUN Test_Normal_10S
--- PASS: Test_Normal_10S (10.00s)
	labs29_test.go:174: 28177
=== RUN Test_Gzip_10S_Level1
--- PASS: Test_Gzip_10S_Level1 (10.00s)
	labs29_test.go:190: 31500
=== RUN Test_Gzip_10S_Level5
--- PASS: Test_Gzip_10S_Level5 (10.00s)
	labs29_test.go:206: 26868
=== RUN Test_Gzip_10S_Level9
--- PASS: Test_Gzip_10S_Level9 (10.00s)
	labs29_test.go:222: 26867
=== RUN Test_Snappy2_10S
--- PASS: Test_Snappy2_10S (10.00s)
	labs29_test.go:237: 22354
PASS
ok  	github.com/idada/go-labs/labs29	50.795s

dada-imac:labs29 dada$ ls -la json.*
-rwxr-xr-x@ 1 dada  staff    46063  4 21 19:40 json.gzip1
-rwxr-xr-x@ 1 dada  staff    25765  4 21 19:40 json.gzip5
-rwxr-xr-x@ 1 dada  staff    18445  4 21 19:40 json.gzip9
-rwxr-xr-x@ 1 dada  staff  4180000  4 21 19:40 json.normal
-rwxr-xr-x@ 1 dada  staff   408426  4 21 19:40 json.snappy
```

实验结论：实验证明各种流压缩算法的吞吐量都高于json序列化过程，所以没必要用太高级压缩算法，可以用最高压缩比gzip