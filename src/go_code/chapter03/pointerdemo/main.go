package main

import "fmt"

// 指针类型
func main() {

	//	基本数据类型的内存布局
	var i int = 10
	//	i的地址是什么，&i
	fmt.Println("i的地址=", &i)

	//1.ptr 是一个指针变量
	//2.ptr 的类型是 *int
	//3.ptr 本身的值&i
	var ptr *int = &i
	fmt.Printf("ptr=%v\n", ptr)
	fmt.Printf("ptr 的地址=%v\n", &ptr)
	fmt.Printf("ptr 指向的值=%v\n", *ptr)
}
