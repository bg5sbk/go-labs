package labs22

import "testing"

import "labs22/mylib"

func Test_Go_Call_C(t *testing.T) {
	mylib.Call(1)
}

func Benchmark_Go_Call_C(b *testing.B) {
	mylib.Call(b.N)
}
