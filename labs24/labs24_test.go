package labs24

import "testing"

func Benchmark_UseBinaryWrite1(b *testing.B) {
	w := &DummyWriter{}
	for i := 0; i < b.N; i++ {
		UseBinaryWrite1(w, int32(i))
	}
}

func Benchmark_UseBinaryWrite2(b *testing.B) {
	w := &DummyWriter{}
	for i := 0; i < b.N; i++ {
		UseBinaryWrite2(w, i)
	}
}

func Benchmark_UseHardcode(b *testing.B) {
	w := &DummyWriter{}
	for i := 0; i < b.N; i++ {
		UseHardcode(w, int32(i))
	}
}
