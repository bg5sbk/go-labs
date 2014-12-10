package labs21

import "testing"

func Benchmark_Go_Call_C(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cfunc()
	}
}

func Benchmark_GO_Call_GO(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gofunc()
	}
}
