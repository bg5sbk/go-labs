package labs12

/*
#cgo LDFLAGS: -ljemalloc
#include <stdlib.h>
#include <jemalloc/jemalloc.h>
*/
import "C"
import "unsafe"

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

type XStruct struct {
	xyz []*BigStruct
}

var sizeof_XStruct = C.size_t(unsafe.Sizeof(XStruct{}))
var sizeof_BigStruct = C.size_t(unsafe.Sizeof(BigStruct{}))

func NewBigStruct() *BigStruct {
	return (*BigStruct)(C.calloc(1, sizeof_BigStruct))
}

func NewXStruct() *XStruct {
	var r = (*XStruct)(C.calloc(1, sizeof_XStruct))
	r.xyz = make([]*BigStruct, 100)
	return r
}

func AllocAndFree() {
	var ptr = C.malloc(sizeof_BigStruct)

	C.free(ptr)
}

func JeAllocAndFree() {
	var ptr = C.je_malloc(sizeof_BigStruct)

	C.je_free(ptr)
}
