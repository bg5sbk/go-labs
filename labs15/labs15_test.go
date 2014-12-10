package labs15

import "runtime"
import "testing"

func Benchmark_Fetch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		temp := make([]MyTable, 0)

		FetchMyTable(func(item MyTable) {
			temp = append(temp, item)
		})
	}

	runtime.GC()
}

func Benchmark_Lookup(b *testing.B) {
	for i := 0; i < b.N; i++ {
		temp := make([]*MyTable, 0)

		for j := 0; j < len(myTable); j++ {
			temp = append(temp, LookupMyTable(j))
		}
	}

	runtime.GC()
}
