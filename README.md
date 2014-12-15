一些杂七杂八脑洞大开的实验：

| 实验 | 描述 |
| ------ | ------ |
| labs01 | 测试类型判断和类型转换的效率 |
| labs02 | 测试值传参和指针传参的效率 |
| labs03 | 测试对象创建的效率 |
| labs04 | 测试range循环和for循环，以及结构体循环和指针循环的性能区别 |
| labs05 | 测试整数和浮点数的运算效率 |
| labs06 | 测试小数据量循环和map取数据以及硬编码取数据的性能消耗 |
| labs07 | 测试几种链表遍历查找的性能差异 |
| labs08 | 同labs07，区别是用slice替代链表结构 |
| labs09 | 测试匿名函数和普通函数的调用消耗 |
| labs10 | 尝试绕过gc扫描 |
| labs11 | 测试json序列化支持的类型 |
| labs12 | 测试jsmalloc和malloc在go程序中是否有性能差别 |
| labs13 | 测试不同数据结构的对象数量差别 |
| labs14 | 测试用cgo复制go的内存数据和还原数据 |
| labs15 | 测试unsafe伪造slice数据结构做内存表的查询效率 |
| labs16 | 测试平方根算法的效率 |
| labs17 | 尝试优化一段代码 |
| labs18 | 测试switch和回调函数效率差异 |
| labs19 | 测试缓存反射信息对效率的影响 |
| labs20 | 测试Go调用C和C调用Go的效率差异 |
| labs21 | 个测试goc机制调用c的实验 |
| labs22 | 尝试把CGO生成的runtime·cgocall替换为runtime·asmcgocall |
| labs23 | 测试interface{}类型转换判断和布尔值判断效率差异 |
| labs24 | 测试binary.Write和硬编码的效率差异 |
