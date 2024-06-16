package main

import "fmt"

/*
Map 是一种无序的键值对的集合。

Map 最重要的一点是通过 key 来快速检索数据，key 类似于索引，指向数据的值。

Map 是一种集合，所以我们可以像迭代数组和切片那样迭代它。不过，Map 是无序的，遍历 Map 时返回的键值对的顺序是不确定的。

在获取 Map 的值时，如果键不存在，返回该类型的零值，例如 int 类型的零值是 0，string 类型的零值是 ""。

Map 是引用类型，如果将一个 Map 传递给一个函数或赋值给另一个变量，它们都指向同一个底层数据结构，因此对 Map 的修改会影响到所有引用它的变量。
*/
func main() {
	//	map的使用和注意事项
	var a map[string]string
	//	使用map前需要先make,make的作用是给map分配数据空间
	a = make(map[string]string, 10)
	a["no1"] = "李四"
	fmt.Println(a)
}
