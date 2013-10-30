package main

import "os"
import "runtime/pprof"

type BigStruct struct {
	C01 int
	C02 int
	C03 int
	C04 int
	C05 int
	C06 int
	C07 int
	C08 int
	C09 int
	C10 int
	C11 int
	C12 int
	C13 int
	C14 int
	C15 int
	C16 int
	C17 int
	C18 int
	C19 int
	C20 int
	C21 int
	C22 int
	C23 int
	C24 int
	C25 int
	C26 int
	C27 int
	C28 int
	C29 int
	C30 int
}

var objects []BigStruct

func main() {
	objects = make([]BigStruct, 0, 10000)

	for i := 0; i < 10000; i++ {
		objects = append(objects, BigStruct{})
	}

	p := pprof.Lookup("heap")
	p.WriteTo(os.Stdout, 2)
}
