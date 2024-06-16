package main

import "fmt"

// 逻辑运算符的使用
func main() {
	var age int = 40

	if age > 30 && age < 50 {
		fmt.Println("&&1")
	}
	if age > 30 && age < 40 {
		fmt.Println("&&2")
	}
	if age > 30 || age < 50 {
		fmt.Println("||3")
	}
	if age > 30 || age < 40 {
		fmt.Println("||4")
	}
	if age > 30 {
		fmt.Println("!5")
	}
	if !(age > 30) {
		fmt.Println("!6")
	}
}
