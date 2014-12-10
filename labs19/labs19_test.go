package labs19

import "testing"

type MyTable struct {
	Name1 string `max-length:"10"`
	Name2 string `max-length:"10"`
	Name3 string `max-length:"10"`
	Name4 string `max-length:"10"`
	Name5 string `max-length:"10"`
	Name6 string `max-length:"10"`
	Name7 string `max-length:"10"`
}

func Test_Verify(t *testing.T) {
	x := new(MyTable)
	x.Name7 = "012"

	if Verify(x) != nil {
		t.FailNow()
	}

	x.Name7 = "0123456789A"

	if Verify(x) == nil {
		t.FailNow()
	}
}

func Test_FastVerify(t *testing.T) {
	x := new(MyTable)
	x.Name7 = "012"

	if FastVerify(x) != nil {
		t.FailNow()
	}

	x.Name7 = "0123456789A"

	if FastVerify(x) == nil {
		t.FailNow()
	}
}

func Test_VeryFastVerify(t *testing.T) {
	x := new(MyTable)
	x.Name7 = "012"

	if VeryFastVerify(x) != nil {
		t.FailNow()
	}

	x.Name7 = "0123456789A"

	if VeryFastVerify(x) == nil {
		t.FailNow()
	}
}

func Benchmark_________Verify(b *testing.B) {
	x := new(MyTable)

	for i := 0; i < b.N; i++ {
		Verify(x)
	}
}

func Benchmark_____FastVerify(b *testing.B) {
	x := new(MyTable)

	for i := 0; i < b.N; i++ {
		FastVerify(x)
	}
}

func Benchmark_VeryFastVerify(b *testing.B) {
	x := new(MyTable)

	for i := 0; i < b.N; i++ {
		VeryFastVerify(x)
	}
}
