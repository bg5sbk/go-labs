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

func Benchmark_VarFuncCall1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		VarFunc := func(i int) int {
			return i + 1
		}

		VarFunc(i)
	}
}

func Benchmark_VarFuncCall2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		VarFunc := func() int {
			return i + 1
		}

		VarFunc()
	}
}
