package labs18

import "encoding/binary"

type MyObject struct {
	n        int
	head     []byte
	callback func([]byte) int
	bo       binary.ByteOrder
}

func New() *MyObject {
	return &MyObject{
		n:    4,
		head: make([]byte, 4),
		bo:   binary.LittleEndian,
	}
}

func NewCallback() *MyObject {
	return &MyObject{
		n:    4,
		head: make([]byte, 4),
		callback: func(head []byte) int {
			return int(binary.LittleEndian.Uint32(head)) + 10
		},
	}
}

func (obj *MyObject) UseSwitch() int {
	size := 0

	switch obj.n {
	case 1:
		size = int(obj.head[0])
	case 2:
		size = int(obj.bo.Uint16(obj.head))
	case 4:
		size = int(obj.bo.Uint32(obj.head))
	case 8:
		size = int(obj.bo.Uint64(obj.head))
	default:
		panic("unsupported packet head size")
	}

	return size + 10
}

func (obj *MyObject) UseCallback() int {
	return obj.callback(obj.head)
}
