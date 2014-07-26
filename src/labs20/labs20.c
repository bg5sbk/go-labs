#include "labs20.h"

void go_call_c() {

}

void c_call_go(int n) {
	int i;

	for (i = 0; i < n; i ++) {
		go_func();
	}
}
