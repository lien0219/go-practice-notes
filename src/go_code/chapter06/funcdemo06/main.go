package main

import "fmt"

var (
	//Fun1就是一个全局匿名函数（首字母必须大写否则无用）
	Fun1 = func(n1 int, n2 int) int {
		return n1 * n2
	}
)

// 匿名函数的使用演示
func main() {
	//	求两数之和
	res1 := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)
	fmt.Println("res1=", res1)

	//	将匿名函数func (n1 int, n2 int) int 赋给a变量
	//	则a的数据类型就是函数类型，此时，可以通过a完成调用
	a := func(n1 int, n2 int) int {
		return n1 - n2
	}

	res2 := a(10, 30)
	fmt.Println("res2=", res2)
	res3 := a(90, 30)
	fmt.Println("res3=", res3)

	//	全局匿名函数的使用
	res4 := Fun1(4, 9)
	fmt.Println("res4=", res4)
}
