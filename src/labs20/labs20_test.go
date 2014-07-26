package labs20

import "testing"

func Benchmark_Go_Call_C(b *testing.B) {
	do_go_call_c(b.N)
}

func Benchmark_C_Call_GO(b *testing.B) {
	do_c_call_go(b.N)
}
