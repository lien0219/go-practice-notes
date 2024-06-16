package main

import "fmt"

// 变量作用域
func test() (int, string) {
	age := 10
	Name := "李四"
	return age, Name
}

func main() {
	fmt.Println(test())
}
