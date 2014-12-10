package labs06

import "testing"

type BigStruct struct {
	C01 int
	C02 int
	C03 int
	C04 int
	C05 int
	C06 int
	C07 int
	C08 int
	C09 int
	C10 int
	C11 int
	C12 int
	C13 int
	C14 int
	C15 int
	C16 int
	C17 int
	C18 int
	C19 int
	C20 int
	C21 int
	C22 int
	C23 int
	C24 int
	C25 int
	C26 int
	C27 int
	C28 int
	C29 int
	C30 int
}

func Loop1(a []BigStruct) int {
	for i := 0; i < len(a); i++ {
		if a[i].C30 == 3 {
			return i
		}
	}

	return -1
}

func Loop2(a []BigStruct) int {
	for i := len(a) - 1; i >= 0; i-- {
		if a[i].C30 == 1 {
			return i
		}
	}

	return -1
}

func Loop3(a map[int]BigStruct) int {
	return a[2].C30
}

func Loop4(a []*BigStruct) int {
	for i, x := range a {
		if x.C30 == 3 {
			return i
		}
	}

	return -1
}

func Loop5(a []BigStruct) int {
	switch {
	case a[0].C01 == 3:
		return 0
	case a[1].C01 == 3:
		return 1
	case a[2].C01 == 3:
		return 2
	}

	return -1
}

func Benchmark_Loop1(b *testing.B) {
	var a = make([]BigStruct, 3)

	a[0].C30 = 1
	a[1].C30 = 2
	a[2].C30 = 3

	for i := 0; i < b.N; i++ {
		Loop1(a)
	}
}

func Benchmark_Loop2(b *testing.B) {
	var a = make([]BigStruct, 3)

	a[0].C30 = 1
	a[1].C30 = 2
	a[2].C30 = 3

	for i := 0; i < b.N; i++ {
		Loop2(a)
	}
}

func Benchmark_Loop3(b *testing.B) {
	var a = make(map[int]BigStruct, 3)

	a[0] = BigStruct{C30: 1}
	a[1] = BigStruct{C30: 2}
	a[2] = BigStruct{C30: 3}

	for i := 0; i < b.N; i++ {
		Loop3(a)
	}
}

func Benchmark_Loop4(b *testing.B) {
	var a = make([]*BigStruct, 3)

	a[0] = &BigStruct{C30: 1}
	a[1] = &BigStruct{C30: 2}
	a[2] = &BigStruct{C30: 3}

	for i := 0; i < b.N; i++ {
		Loop4(a)
	}
}

func Benchmark_Loop5(b *testing.B) {
	var a = make([]BigStruct, 3)

	a[0].C30 = 1
	a[1].C30 = 2
	a[2].C30 = 3

	for i := 0; i < b.N; i++ {
		Loop5(a)
	}
}
