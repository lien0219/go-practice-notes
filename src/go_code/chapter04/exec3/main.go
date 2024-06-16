package main

import "fmt"

// 数组练习
func main() {

	//	1.两数最大值
	var n1 int = 10
	var n2 int = 40
	var max int

	if n1 > n2 {
		max = n1
	} else {
		max = n2
	}
	fmt.Println("max=", max)

	//	2.三个数最大值
	var n3 int = 45

	if n3 > max {
		max = n3
	}
	fmt.Println("三个数最大值是=", max)
}
