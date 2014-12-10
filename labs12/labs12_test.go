package labs12

import "testing"

func Benchmark_AllocAndFree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AllocAndFree()
	}
}

func Benchmark_JeAllocAndFree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		JeAllocAndFree()
	}
}

func Test_Nil(t *testing.T) {
	var a = NewBigStruct()

	if a.next != nil {
		t.Fail()
	}

	var b = NewXStruct()

	for i := 0; i < 100; i++ {
		if b.xyz[i] != nil {
			t.Fail()
		}
	}
}
