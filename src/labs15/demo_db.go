package labs15

/*
#include <stdlib.h>
*/
import "C"
import "unsafe"
import "reflect"

type MyTable struct {
	C01   int
	C02   int
	C03   int
	C04   int
	C05   int
	C06   int
	C07   int
	C08   int
	C09   int
	C10   int
	Child *ChildData
}

type ChildData struct {
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
}

var myTable []MyTable

func init() {
	size := 100000
	head := (*reflect.SliceHeader)(unsafe.Pointer(&myTable))
	head.Data = uintptr(C.calloc(C.size_t(size), C.size_t(unsafe.Sizeof(MyTable{}))))
	head.Cap = size
	head.Len = size

	sizeofChildData := C.size_t(unsafe.Sizeof(ChildData{}))

	for i := 0; i < size; i++ {
		myTable[i].C01 = i
		myTable[i].Child = (*ChildData)(C.calloc(1, sizeofChildData))
		myTable[i].Child.C01 = i
	}
}

func FetchMyTable(callback func(item MyTable)) {
	for i := 0; i < len(myTable); i++ {
		callback(myTable[i])
	}
}

func LookupMyTable(i int) *MyTable {
	temp := myTable[i]
	return &temp
}
