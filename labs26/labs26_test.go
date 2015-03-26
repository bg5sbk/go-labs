package labs26

import (
	"reflect"
	"testing"
)

func MyFunc(a, b, c int) int {
	return a + b + c
}

func Benchmark_NormalFuncCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MyFunc(i, i, i)
	}
}

func Benchmark_ReflectFuncCall(b *testing.B) {
	f := reflect.ValueOf(MyFunc)
	for i := 0; i < b.N; i++ {
		f.Call([]reflect.Value{
			reflect.ValueOf(i),
			reflect.ValueOf(i),
			reflect.ValueOf(i),
		})
	}
}
