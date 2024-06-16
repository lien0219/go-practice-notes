package main

import (
	"fmt"
	"strings"
)

func AddUpper() func(int) int {
	var n int = 10
	return func(x int) int {
		n += x
		return n
	}
}

/*
编写一个函数，可以接收一个文件后缀名，并返回一个闭包
调用闭包可以传入一个文件名，如果该文件没有指定的后缀名则返回文件名

strings.HasSuffix  该函数可以判断某个字符串是否有指定的后缀
*/
func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

// 闭包演示
func main() {
	f := AddUpper()
	fmt.Println(f(1))

	//	测试makeSuffix
	makeSuffix := makeSuffix(".jpg")
	fmt.Println("文件名处理后：", makeSuffix("winter"))
}
