package labs23

import (
	"testing"
)

func Benchmark_UseInterface1(b *testing.B) {
	obj := &MyObject{Conn: new(BufferConn)}
	for i := 0; i < b.N; i++ {
		obj.UseInterface(i)
	}
}

func Benchmark_UseInterface2(b *testing.B) {
	obj := &MyObject{Conn: new(NormalConn)}
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
