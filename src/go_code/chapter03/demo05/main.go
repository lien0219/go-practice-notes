package main

import "fmt"

// golang中+的使用
func main() {
	//数字类型加
	var i = 1
	var j = 2
	var r = i + j
	fmt.Println("r=", r)

	//字符串类型拼接
	var str1 = "hello"
	var str2 = "world"
	var res = str1 + str2
	fmt.Println("res=", res)
}
