package main

import "fmt"

func main() {
	var num int
	//	常量声明的时候必须赋值且不能修改
	//	常量只能修饰bool、数值类型（int,float系列）、string
	const tax int = 0
	fmt.Println(num, tax)

	//	多个常量简洁写法
	const (
		a = 1
		b = 2
	)
	fmt.Println(a, b)
	//	或者这样写
	const (
		a1 = iota
		b1
		c1
	)
	fmt.Println(a1, b1, c1) //0，1，2
}
