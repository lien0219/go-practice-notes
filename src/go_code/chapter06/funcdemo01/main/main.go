package main

import (
	"fmt"
	"go_code/chapter06/funcdemo01/utils"
	//util "go_code/chapter06/funcdemo01/utils"   别名写法，别名在前
)

func main() {
	//输入两个数再输入一个运算符得到结果
	var n1 float64 = 1.2
	var n2 float64 = 2.3
	var operator byte = '+'

	result := utils.Cal(n1, n2, operator)
	fmt.Println("res=", result)
}
