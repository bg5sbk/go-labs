package labs20

/*
#include "labs20.h"
*/
import "C"

//export go_func
func go_func() {

}

func do_go_call_c(n int) {
	for i := 0; i < n; i++ {
		C.go_call_c()
	}
}

func do_c_call_go(n int) {
	C.c_call_go(C.int(n))
}
