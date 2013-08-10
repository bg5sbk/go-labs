package labs07

import "testing"
import "reflect"

var data *BigStruct = &BigStruct{C30: 88888888}

func init() {
	for i := 0; i < 1000; i++ {
		var newData = new(BigStruct)
		newData.C30 = i
		newData.next = data
		data = newData
	}
}

func Loop1() *BigStruct {
	for n := data; n != nil; n = n.next {
		if n.C30 == 88888888 {
			return n
		}
	}

	return nil
}

func Loop2(callback func(*BigStruct) bool) *BigStruct {
	for n := data; n != nil; n = n.next {
		if callback(n) {
			return n
		}
	}

	return nil
}

func Loop3(callback func(BigStruct) bool) *BigStruct {
	for n := data; n != nil; n = n.next {
		if callback(*n) {
			return n
		}
	}

	return nil
}

func Loop4(callback func(*BigStruct) bool) *BigStruct {
	for n := data; n != nil; n = n.next {
		nn := *n
		if callback(&nn) {
			return n
		}
	}

	return nil
}

func Loop5(name string, value interface{}) *BigStruct {
	for n := data; n != nil; n = n.next {
		var v = reflect.ValueOf(n)

		if reflect.DeepEqual(v.Elem().FieldByName(name).Interface(), value) {
			return n
		}
	}

	return nil
}

func Loop6(query *Query) *BigStruct {
	for n := data; n != nil; n = n.next {
		if query.Match(n) {
			var nn = *n
			return &nn
		}
	}
	return nil
}

// 基准测试
//
func Benchmark_Loop1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Loop1()
	}
}

// 测试指针传递进行查找的效率（外部可修改内部数据）
//
func Benchmark_Loop2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Loop2(func(n *BigStruct) bool {
			return n.C30 == 88888888
		})
	}
}

// 测试值传递进行查找的效率（外部不可修改内部数据）
//
func Benchmark_Loop3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Loop3(func(n BigStruct) bool {
			return n.C30 == 88888888
		})
	}
}

// 测试复制数据后传递指针进行查找的效率（外部不可修改内部数据）
//
func Benchmark_Loop4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Loop4(func(n *BigStruct) bool {
			return n.C30 == 88888888
		})
	}
}

// 测试反射取值（自定义查询表达式的可能性）
//
func Benchmark_Loop5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Loop5("C30", 88888888)
	}
}

// 测试指针取值（自定义查询表达式的可能性）
//
func Benchmark_Loop6(b *testing.B) {
	var query = NewQuery("C30", "==", 88888888)

	for i := 0; i < b.N; i++ {
		Loop6(query)
	}
}

func Test_Loop4(t *testing.T) {
	var a = new(BigStruct)
	a.C30 = 100

	var b = *a
	var c = &b

	c.C30 = 200

	if a.C30 == c.C30 {
		t.Fail()
	}

	if b.C30 != c.C30 {
		t.Fail()
	}
}

func Test_Loop5(t *testing.T) {
	if Loop5("C30", 88888888) == nil {
		t.Fail()
	}
}

func Test_Loop6(t *testing.T) {
	var query = NewQuery("C30", "==", 88888888)

	if Loop6(query) == nil {
		t.Fail()
	}
}
