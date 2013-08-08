package labs02

import "testing"

type BigStruct struct {
	C01 uint64
	C02 uint64
	C03 uint64
	C04 uint64
	C05 uint64
	C06 uint64
	C07 uint64
	C08 uint64
	C09 uint64
	C10 uint64
	C11 uint64
	C12 uint64
	C13 uint64
	C14 uint64
	C15 uint64
	C16 uint64
	C17 uint64
	C18 uint64
	C19 uint64
	C20 uint64
	C21 uint64
	C22 uint64
	C23 uint64
	C24 uint64
	C25 uint64
	C26 uint64
	C27 uint64
	C28 uint64
	C29 uint64
	C30 uint64
}

func Invoke1(a *BigStruct) uint64 {
	return a.C30
}

func Invoke2(a BigStruct) uint64 {
	return a.C30
}

func Benchmark_Invoke1(b *testing.B) {
	var a = new(BigStruct)

	for i := 0; i < b.N; i++ {
		Invoke1(a)
	}
}

func Benchmark_Invoke2(b *testing.B) {
	var a = BigStruct{}

	for i := 0; i < b.N; i++ {
		Invoke2(a)
	}
}
