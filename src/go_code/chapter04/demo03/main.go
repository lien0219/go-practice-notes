package main

import "fmt"

// 关系运算符的使用
func main() {
	var n1 int = 9
	var n2 int = 8
	fmt.Println(n1 == n2) //false
	fmt.Println(n1 != n2) //true
	fmt.Println(n1 > n2)  //true
	fmt.Println(n1 >= n2) //true
	fmt.Println(n1 < n2)  //false
	fmt.Println(n1 <= n2) //false

	flag := n1 > n2

	fmt.Println("flag=", flag)
}
