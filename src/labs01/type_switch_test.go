package labs01

import "testing"

type InterfaceA interface {
	AA()
}

type InterfaceB interface {
	BB()
}

type A struct {
	v int
}

func (a *A) AA() {
	a.v += 1
}

type B struct {
	v int
}

func (b *B) BB() {
	b.v += 1
}

func TypeSwitch(v interface{}) {
	switch v.(type) {
	case InterfaceA:
		v.(InterfaceA).AA()
	case InterfaceB:
		v.(InterfaceB).BB()
	}
}

func NormalSwitch(a *A) {
	a.AA()
}

func InterfaceSwitch(v interface{}) {
	v.(InterfaceA).AA()
}

func Benchmark_TypeSwitch(b *testing.B) {
	var a = new(A)

	for i := 0; i < b.N; i++ {
		TypeSwitch(a)
	}
}

func Benchmark_NormalSwitch(b *testing.B) {
	var a = new(A)

	for i := 0; i < b.N; i++ {
		NormalSwitch(a)
	}
}

func Benchmark_InterfaceSwitch(b *testing.B) {
	var a = new(A)

	for i := 0; i < b.N; i++ {
		InterfaceSwitch(a)
	}
}
