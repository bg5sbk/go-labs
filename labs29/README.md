测试不同压缩算法压缩json数据的压缩比和压缩效率。

实验结果：

```
dada-imac:labs29 dada$ go test -v
=== RUN Test_Normal
--- PASS: Test_Normal (0.16s)
=== RUN Test_Gzip_Level1
--- PASS: Test_Gzip_Level1 (0.13s)
=== RUN Test_Gzip_Level5
--- PASS: Test_Gzip_Level5 (0.15s)
=== RUN Test_Gzip_Level9
--- PASS: Test_Gzip_Level9 (0.15s)
=== RUN Test_Snappy
--- PASS: Test_Snappy (0.14s)
=== RUN Test_Normal_10S
--- PASS: Test_Normal_10S (10.00s)
	labs29_test.go:155: 28 MB/S
=== RUN Test_Gzip_10S_Level1
--- PASS: Test_Gzip_10S_Level1 (10.00s)
	labs29_test.go:173: 29 MB/S
=== RUN Test_Gzip_10S_Level5
--- PASS: Test_Gzip_10S_Level5 (10.00s)
	labs29_test.go:191: 27 MB/S
=== RUN Test_Gzip_10S_Level9
--- PASS: Test_Gzip_10S_Level9 (10.00s)
	labs29_test.go:208: 25 MB/S
=== RUN Test_Snappy2_10S
--- PASS: Test_Snappy2_10S (10.00s)
	labs29_test.go:224: 21 MB/S
PASS
ok  	github.com/idada/go-labs/labs29	50.738s

dada-imac:labs29 dada$ ls -lah
total 9192
drwxr-xr-x@ 10 dada  staff   340B  4 22 01:54 .
drwxr-xr-x@ 34 dada  staff   1.1K  4 22 00:05 ..
-rw-r--r--@  1 dada  staff   1.6K  4 22 09:24 README.md
-rwxr-xr-x@  1 dada  staff    45K  4 22 09:22 json.gzip1
-rwxr-xr-x@  1 dada  staff    25K  4 22 09:22 json.gzip5
-rwxr-xr-x@  1 dada  staff    18K  4 22 09:22 json.gzip9
-rwxr-xr-x@  1 dada  staff   4.0M  4 22 09:22 json.normal
-rwxr-xr-x@  1 dada  staff   399K  4 22 09:22 json.snappy
-rw-r--r--@  1 dada  staff   4.0K  4 22 09:22 labs29_test.go
```

实验结论：实验证明各种流压缩算法的吞吐量都高于json序列化过程，所以没必要用太高级压缩算法，可以用最高压缩比gzip