package labs05

import "testing"

func Benchmark_IntAdd(b *testing.B) {
	var a = 0

	for i := 0; i < b.N; i++ {
		a += 1
	}
}

func Benchmark_Int8Add(b *testing.B) {
	var a int8 = 0

	for i := 0; i < b.N; i++ {
		a += 1
	}
}

func Benchmark_Int16Add(b *testing.B) {
	var a int8 = 0

	for i := 0; i < b.N; i++ {
		a += 1
	}
}

func Benchmark_Int32Add(b *testing.B) {
	var a int32 = 0

	for i := 0; i < b.N; i++ {
		a += 1
	}
}

func Benchmark_Int64Add(b *testing.B) {
	var a int64 = 0

	for i := 0; i < b.N; i++ {
		a += 1
	}
}

func Benchmark_Float32Add(b *testing.B) {
	var a float32 = 0.1

	for i := 0; i < b.N; i++ {
		a += 1.0
	}
}

func Benchmark_Float64Add(b *testing.B) {
	var a float64 = 0.1

	for i := 0; i < b.N; i++ {
		a += 1.0
	}
}

func Benchmark_IntSub(b *testing.B) {
	var a = 0x7FFFFFFFFF

	for i := 0; i < b.N; i++ {
		a -= 1
	}
}

func Benchmark_Int8Sub(b *testing.B) {
	var a int8 = 0x7F

	for i := 0; i < b.N; i++ {
		a -= 1
	}
}

func Benchmark_Int16Sub(b *testing.B) {
	var a int16 = 0x7FFF

	for i := 0; i < b.N; i++ {
		a -= 1
	}
}

func Benchmark_Int32Sub(b *testing.B) {
	var a int32 = 0x7FFFFFFF

	for i := 0; i < b.N; i++ {
		a -= 1
	}
}

func Benchmark_Int64Sub(b *testing.B) {
	var a int64 = 0x7FFFFFFFFF

	for i := 0; i < b.N; i++ {
		a -= 1
	}
}

func Benchmark_Float32Sub(b *testing.B) {
	var a = float32(0x7FFFFFFF)

	for i := 0; i < b.N; i++ {
		a -= 1.0
	}
}

func Benchmark_Float64Sub(b *testing.B) {
	var a = float64(0xFFFFFFFFFF)

	for i := 0; i < b.N; i++ {
		a -= 1.0
	}
}

func Benchmark_IntMul(b *testing.B) {
	var a = 1

	for i := 0; i < b.N; i++ {
		a *= 3
	}
}

func Benchmark_Int8Mul(b *testing.B) {
	var a int8 = 1

	for i := 0; i < b.N; i++ {
		a *= 3
	}
}

func Benchmark_Int16Mul(b *testing.B) {
	var a int16 = 1

	for i := 0; i < b.N; i++ {
		a *= 3
	}
}

func Benchmark_Int32Mul(b *testing.B) {
	var a int32 = 1

	for i := 0; i < b.N; i++ {
		a *= 3
	}
}

func Benchmark_Int64Mul(b *testing.B) {
	var a int64 = 1

	for i := 0; i < b.N; i++ {
		a *= 3
	}
}

func Benchmark_Float32Mul(b *testing.B) {
	var a float32 = 1.0

	for i := 0; i < b.N; i++ {
		a *= 1.5
	}
}

func Benchmark_Float64Mul(b *testing.B) {
	var a float64 = 1.0

	for i := 0; i < b.N; i++ {
		a *= 1.5
	}
}

func Benchmark_IntDiv(b *testing.B) {
	var a = 0x7FFFFFFFFF

	for i := 0; i < b.N; i++ {
		a /= 3
	}
}

func Benchmark_Int8Div(b *testing.B) {
	var a int8 = 0x7F

	for i := 0; i < b.N; i++ {
		a /= 3
	}
}

func Benchmark_Int16Div(b *testing.B) {
	var a int16 = 0x7FFF

	for i := 0; i < b.N; i++ {
		a /= 3
	}
}

func Benchmark_Int32Div(b *testing.B) {
	var a int32 = 0x7FFFFFFF

	for i := 0; i < b.N; i++ {
		a /= 3
	}
}

func Benchmark_Int64Div(b *testing.B) {
	var a int64 = 0x7FFFFFFFFF

	for i := 0; i < b.N; i++ {
		a /= 3
	}
}

func Benchmark_Float32Div(b *testing.B) {
	var a = float32(0x7FFFFFFF)

	for i := 0; i < b.N; i++ {
		a /= 1.5
	}
}

func Benchmark_Float64Div(b *testing.B) {
	var a = float64(0x7FFFFFFFFF)

	for i := 0; i < b.N; i++ {
		a /= 1.5
	}
}
