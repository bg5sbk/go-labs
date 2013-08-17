package labs10

import (
	"unsafe"
)

func xnew(size int) unsafe.Pointer

func xfree(p unsafe.Pointer)
