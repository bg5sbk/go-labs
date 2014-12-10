package labs18

import "testing"

func Benchmark_UseSwitch(b *testing.B) {
	obj := New()
	for i := 0; i < b.N; i++ {
		obj.UseSwitch()
	}
}

func Benchmark_UseCallback(b *testing.B) {
	obj := NewCallback()
	for i := 0; i < b.N; i++ {
		obj.UseCallback()
	}
}
