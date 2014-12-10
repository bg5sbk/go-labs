package main

/*
#include <stdlib.h>
#include <string.h>
*/
import "C"
import "reflect"
import "unsafe"
import "strconv"

type SomeStruct struct {
	Id   int
	Name string
}

var someData []SomeStruct

func init() {
	data := new(reflect.SliceHeader)
	data.Cap = 10
	data.Len = 10
	data.Data = uintptr(C.calloc(1, C.size_t(unsafe.Sizeof(SomeStruct{}))))

	someData = *(*[]SomeStruct)(unsafe.Pointer(data))

	for i := 0; i < len(someData); i++ {
		someData[i] = SomeStruct{
			Id:   i,
			Name: "你好" + strconv.Itoa(i),
		}
	}
}

func takeData(i int) SomeStruct {
	return someData[i]
}

func fetchData(callback func(s SomeStruct)) {
	for i := 0; i < len(someData); i++ {
		callback(someData[i])
	}
}

func main() {
	for i := 0; i < len(someData); i++ {
		println(takeData(i).Name)
	}

	var data []SomeStruct

	fetchData(func(item SomeStruct) {
		data = append(data, item)
	})

	for i := 0; i < len(data); i++ {
		println(data[i].Name)
	}
}
