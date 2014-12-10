package labs10

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

var SizeOfBigStruct = int(unsafe.Sizeof(BigStruct{}))

func NewBigStruct() *BigStruct {
	return (*BigStruct)(xnew(SizeOfBigStruct))
}

func FreeBigStruct(p *BigStruct) {
	xfree(unsafe.Pointer(p))
}
