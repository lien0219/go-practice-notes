package main

import "fmt"

// 位运算的测试
func main() {
	fmt.Println(2 & 3)  //2
	fmt.Println(2 | 3)  //3
	fmt.Println(2 ^ 3)  //1
	fmt.Println(-2 ^ 2) //-4

	//	位移运算
	a := 1 >> 2 //0
	c := 1 << 2 //4
	fmt.Println("a=", a, "c=", c)
}
