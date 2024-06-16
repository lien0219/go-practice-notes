package main

import "fmt"

// 指针的 & 和 *运算符的使用
func main() {
	a := 100
	fmt.Println("a 的地址=", &a)

	var ptr *int = &a
	fmt.Println("ptr 指向的值是=", *ptr)
}
