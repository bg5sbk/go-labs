package mylib

/*
void cfunc() {
}
*/
import "C"

func Call(n int) {
	for i := 0; i < n; i++ {
		C.cfunc()
	}
}
