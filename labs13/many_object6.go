package main

import "os"
import "runtime/pprof"

type BigStruct struct {
	C01 []int
	C02 []int
	C03 []int
	C04 []int
	C05 []int
	C06 []int
	C07 []int
	C08 []int
	C09 []int
	C10 []int
	C11 []int
	C12 []int
	C13 []int
	C14 []int
	C15 []int
	C16 []int
	C17 []int
	C18 []int
	C19 []int
	C20 []int
	C21 []int
	C22 []int
	C23 []int
	C24 []int
	C25 []int
	C26 []int
	C27 []int
	C28 []int
	C29 []int
	C30 []int
}

func main() {
	data := make([]BigStruct, 1000)
	for i := 0; i < len(data); i++ {
		data[i] = BigStruct{
			C01: make([]int, 0, 1),
			C02: make([]int, 0, 1),
			C03: make([]int, 0, 1),
			C04: make([]int, 0, 1),
			C05: make([]int, 0, 1),
			C06: make([]int, 0, 1),
			C07: make([]int, 0, 1),
			C08: make([]int, 0, 1),
			C09: make([]int, 0, 1),
			C10: make([]int, 0, 1),
			C11: make([]int, 0, 1),
			C12: make([]int, 0, 1),
			C13: make([]int, 0, 1),
			C14: make([]int, 0, 1),
			C15: make([]int, 0, 1),
			C16: make([]int, 0, 1),
			C17: make([]int, 0, 1),
			C18: make([]int, 0, 1),
			C19: make([]int, 0, 1),
			C20: make([]int, 0, 1),
			C21: make([]int, 0, 1),
			C22: make([]int, 0, 1),
			C23: make([]int, 0, 1),
			C24: make([]int, 0, 1),
			C25: make([]int, 0, 1),
			C26: make([]int, 0, 1),
			C27: make([]int, 0, 1),
			C28: make([]int, 0, 1),
			C29: make([]int, 0, 1),
			C30: make([]int, 0, 1),
		}
	}
	p := pprof.Lookup("heap")
	p.WriteTo(os.Stdout, 2)
}
