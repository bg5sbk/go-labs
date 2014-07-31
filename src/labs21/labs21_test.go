package labs21

import "testing"

func Benchmark_Go_Call_C1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cfunc1()
	}
}

/*
func Benchmark_Go_Call_C2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cfunc2()
	}
}
*/
