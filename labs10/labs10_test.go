package labs10

import "testing"

func Test_NewBigStruct(t *testing.T) {
	bs := NewBigStruct()

	FreeBigStruct(bs)
}
