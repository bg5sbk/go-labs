package labs09

import "testing"

func NormalFunc(i int) int {
	return i + 1
}

func Benchmark_NormalFuncCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NormalFunc(i)
	}
}

func Benchmark_VarFuncCall(b *testing.B) {
	VarFunc := func(i int) int {
		return i + 1
	}

	for i := 0; i < b.N; i++ {
		VarFunc(i)
	}
}
