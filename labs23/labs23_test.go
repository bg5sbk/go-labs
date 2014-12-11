package labs23

import (
	"testing"
)

func Benchmark_UseInterface(b *testing.B) {
	obj := &MyObject{Conn: new(BufferConn)}
	for i := 0; i < b.N; i++ {
		obj.UseInterface(i)
	}
}

func Benchmark_UseBoolean(b *testing.B) {
	obj := &MyObject{UseBufferConn: true, Conn: new(BufferConn)}
	for i := 0; i < b.N; i++ {
		obj.UseBoolean(i)
	}
}
