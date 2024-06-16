package main

import "fmt"

// 练习面试题
func main() {
	//有两个变量，a和b，要求将其进行交换，但是不允许使用中间变量，最终打印结果
	var a int = 10
	var b int = 20
	//方法1
	//a = a + b
	//b = a - b
	//a = a - b

	//方法2
	a, b = b, a
	fmt.Printf("a=%v b=%v", a, b)
}
