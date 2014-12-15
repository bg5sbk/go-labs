package labs24

import (
	"encoding/binary"
	"io"
)

type DummyWriter struct {
}

func (w *DummyWriter) Write(b []byte) (int, error) {
	return len(b), nil
}

func UseBinaryWrite1(w io.Writer, i int32) {
	binary.Write(w, binary.LittleEndian, i)
}

func UseBinaryWrite2(w io.Writer, i int) {
	binary.Write(w, binary.LittleEndian, i)
}

func UseHardcode(w io.Writer, i int32) {
	w.Write([]byte{
		byte(i),
		byte(i >> 8),
		byte(i >> 16),
		byte(i >> 24),
	})
}
