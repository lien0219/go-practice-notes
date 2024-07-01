package main

import "fmt"

// 测试函数
func addUpper(n int) int {
	res := 0
	for i := 0; i <= n; i++ {
		res += i
	}
	return res
}

func main() {
	res := addUpper(10)
	if res != 55 {
		fmt.Printf("addUpper错误 返回值：%v 期望值：%v\n", res, 55)
	} else {
		fmt.Printf("addUpper正确 返回值：%v 期望值：%v\n", res, 55)
	}
}
