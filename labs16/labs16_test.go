package labs16

import (
	"math"
	"testing"
	"unsafe"
)

func Benchmark_Method1(b *testing.B) {
	for m := 0; m < b.N; m++ {
		var x = float32(10012312312)
		var z = 0.5 * x
		for i := 0; i < 1000; i++ {
			z -= (z*z - x) / (2 * x)
		}
	}
}

func Benchmark_Method2(b *testing.B) {
	for m := 0; m < b.N; m++ {
		var x = float32(10012312312)
		var xhalf = 0.5 * x

		i := *(*int32)(unsafe.Pointer(&x))
		i = 0x5f3759df - (i >> 1)
		x = *(*float32)(unsafe.Pointer(&i))
		x = x * (1.5 - xhalf*x*x)
		x = x * (1.5 - xhalf*x*x)
		x = 1 / x
	}
}

func Benchmark_Method3(b *testing.B) {
	for m := 0; m < b.N; m++ {
		math.Sqrt(10012312312)
	}
}
