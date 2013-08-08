package labs03

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

func NewStruct1() *BigStruct {
	return new(BigStruct)
}

func NewStruct2() *BigStruct {
	return &BigStruct{}
}

func NewStruct3() *BigStruct {
	var r = new(BigStruct)
	*r = BigStruct{}
	return r
}

func NewStruct4(r *BigStruct) {
	*r = BigStruct{}
}

func NewStruct5(r *BigStruct) {
	r.C01 = 0
	r.C02 = 0
	r.C03 = 0
	r.C04 = 0
	r.C05 = 0
	r.C06 = 0
	r.C07 = 0
	r.C08 = 0
	r.C09 = 0
	r.C10 = 0
	r.C11 = 0
	r.C12 = 0
	r.C13 = 0
	r.C14 = 0
	r.C15 = 0
	r.C16 = 0
	r.C17 = 0
	r.C18 = 0
	r.C19 = 0
	r.C20 = 0
	r.C21 = 0
	r.C22 = 0
	r.C23 = 0
	r.C24 = 0
	r.C25 = 0
	r.C26 = 0
	r.C27 = 0
	r.C28 = 0
	r.C29 = 0
	r.C30 = 0
}

func Benchmark_NewStruct1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewStruct1()
	}
}

func Benchmark_NewStruct2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewStruct2()
	}
}

func Benchmark_NewStruct3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewStruct3()
	}
}

func Benchmark_NewStruct4(b *testing.B) {
	var r = new(BigStruct)

	for i := 0; i < b.N; i++ {
		NewStruct4(r)
	}
}

func Benchmark_NewStruct5(b *testing.B) {
	var r = new(BigStruct)

	for i := 0; i < b.N; i++ {
		NewStruct5(r)
	}
}
