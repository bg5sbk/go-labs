package main

/*
#include <string.h>
*/
import "C"
import "time"
import "reflect"
import "unsafe"

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

func initTestData() []BigStruct {
	data := make([]BigStruct, 1000)
	for i := 0; i < len(data); i++ {
		n := i * 30
		data[i] = BigStruct{
			C01: n + 1,
			C02: n + 2,
			C03: n + 3,
			C04: n + 4,
			C05: n + 5,
			C06: n + 6,
			C07: n + 7,
			C08: n + 8,
			C09: n + 9,
			C10: n + 10,
			C11: n + 11,
			C12: n + 12,
			C13: n + 13,
			C14: n + 14,
			C15: n + 15,
			C16: n + 16,
			C17: n + 17,
			C18: n + 18,
			C19: n + 19,
			C20: n + 20,
			C21: n + 21,
			C22: n + 22,
			C23: n + 23,
			C24: n + 24,
			C25: n + 25,
			C26: n + 26,
			C27: n + 27,
			C28: n + 28,
			C29: n + 29,
			C30: n + 30,
		}
	}
	return data
}

func checkTestData(length int, data []BigStruct) {
	assert(len(data) == length)
	for i := 0; i < len(data); i++ {
		n := i * 30
		assert(data[i].C01 == n+1)
		assert(data[i].C02 == n+2)
		assert(data[i].C03 == n+3)
		assert(data[i].C04 == n+4)
		assert(data[i].C05 == n+5)
		assert(data[i].C06 == n+6)
		assert(data[i].C07 == n+7)
		assert(data[i].C08 == n+8)
		assert(data[i].C09 == n+9)
		assert(data[i].C10 == n+10)
		assert(data[i].C11 == n+11)
		assert(data[i].C12 == n+12)
		assert(data[i].C13 == n+13)
		assert(data[i].C14 == n+14)
		assert(data[i].C15 == n+15)
		assert(data[i].C16 == n+16)
		assert(data[i].C17 == n+17)
		assert(data[i].C18 == n+18)
		assert(data[i].C19 == n+19)
		assert(data[i].C20 == n+20)
		assert(data[i].C21 == n+21)
		assert(data[i].C22 == n+22)
		assert(data[i].C23 == n+23)
		assert(data[i].C24 == n+24)
		assert(data[i].C25 == n+25)
		assert(data[i].C26 == n+26)
		assert(data[i].C27 == n+27)
		assert(data[i].C28 == n+28)
		assert(data[i].C29 == n+29)
		assert(data[i].C30 == n+30)
	}
}

func assert(c bool) {
	if !c {
		panic("assert failed")
	}
}

func main() {
	data1 := initTestData()

	checkTestData(len(data1), data1)

	itemSize := unsafe.Sizeof(BigStruct{})

	csize := C.size_t(itemSize * uintptr(len(data1)))

	t1 := time.Now().Nanosecond()

	mem := C.malloc(csize)

	C.memcpy(mem, unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&data1)).Data), csize)

	t2 := time.Now().Nanosecond()

	println("Go to C:", (t2-t1)/int(time.Microsecond))

	data2 := new(reflect.SliceHeader)
	data2.Cap = len(data1)
	data2.Len = len(data1)
	data2.Data = uintptr(mem)

	checkTestData(len(data1), *(*[]BigStruct)(unsafe.Pointer(data2)))

	t3 := time.Now().Nanosecond()

	data3 := make([]BigStruct, len(data1))
	copy(data3, *(*[]BigStruct)(unsafe.Pointer(data2)))

	t4 := time.Now().Nanosecond()

	checkTestData(len(data1), data3)

	println("C to Go:", (t4-t3)/int(time.Microsecond))
}
