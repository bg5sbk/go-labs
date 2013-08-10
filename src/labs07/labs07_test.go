package labs07

import "testing"

type BigStruct struct {
	next *BigStruct
	C01  int
	C02  int
	C03  int
	C04  int
	C05  int
	C06  int
	C07  int
	C08  int
	C09  int
	C10  int
	C11  int
	C12  int
	C13  int
	C14  int
	C15  int
	C16  int
	C17  int
	C18  int
	C19  int
	C20  int
	C21  int
	C22  int
	C23  int
	C24  int
	C25  int
	C26  int
	C27  int
	C28  int
	C29  int
	C30  int
}

var data *BigStruct

func init() {
	for i := 0; i < 1000; i++ {
		var newData = new(BigStruct)
		newData.C30 = i
		newData.next = data
		data = newData
	}
}

func Loop1() *BigStruct {
	for n := data; n != nil; n = n.next {
		if n.C30 == 0 {
			return n
		}
	}

	return nil
}

func Loop2(callback func(*BigStruct) bool) *BigStruct {
	for n := data; n != nil; n = n.next {
		if callback(n) {
			return n
		}
	}

	return nil
}

func Loop3(callback func(BigStruct) bool) *BigStruct {
	for n := data; n != nil; n = n.next {
		if callback(*n) {
			return n
		}
	}

	return nil
}

func Loop4(callback func(*BigStruct) bool) *BigStruct {
	for n := data; n != nil; n = n.next {
		nn := *n
		if callback(&nn) {
			return n
		}
	}

	return nil
}

func Benchmark_Loop1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Loop1()
	}
}

func Benchmark_Loop2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Loop2(func(n *BigStruct) bool {
			return n.C30 == 0
		})
	}
}

func Benchmark_Loop3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Loop3(func(n BigStruct) bool {
			return n.C30 == 0
		})
	}
}

func Benchmark_Loop4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Loop2(func(n *BigStruct) bool {
			return n.C30 == 0
		})
	}
}

func Test_Loop4(t *testing.T) {
	var a = new(BigStruct)
	a.C30 = 100

	var b = *a
	var c = &b

	c.C30 = 200

	if a.C30 == c.C30 {
		t.Fail()
	}

	if b.C30 != c.C30 {
		t.Fail()
	}
}
