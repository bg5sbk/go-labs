#include "runtime.h"

void* runtime·mallocgc(uintptr size, uint32 flag, int32 dogc, int32 zeroed);

void ·xfree(void *p){
	runtime·free(*(&p));
}

void ·xnew(uintptr size,void *p){
	*(&p) = runtime·mallocgc(size, 1<<2, 0, 0);
}
