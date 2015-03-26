package labs27

import "testing"

func Benchmark_Goid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goid()
	}
}
