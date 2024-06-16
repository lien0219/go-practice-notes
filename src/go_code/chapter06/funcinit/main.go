package main

import (
	"fmt"
	"go_code/chapter06/funcinit/utils"
)

var age = test()

// 方便看到全局变量先被初始化的，
func test() int {
	fmt.Println("test...")
	return 90
}

// init函数，通常可以在init函数中完成初始化的工作
func init() {
	fmt.Println("init...")
}

func main() {
	fmt.Println("main....", age)
	fmt.Println("Age=", utils.Age, "Name=", utils.Name)
}
