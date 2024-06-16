package main

import "fmt"

// 条件语句练习
func main() {
	//   1.编写程序，声明2个int32型变量并赋值。判断两数之和，如果大于等于50，打印“hello world"
	var num1 int32 = 20
	var num2 int32 = 40
	if num1+num2 >= 50 {
		fmt.Println("hello world")
	}

	//	2.编写程序，生命两个float64型变量并赋值，判断第一个数大于10.0，且第二个数小于20.0，打印两数之和
	var n1 float64 = 12.5
	var n2 float64 = 15.6
	if n1 > 10.0 && n2 < 20.0 {
		fmt.Println("和=", (n1 + n2))
	}

	//	3.定义两个变量int32,判断两者的和，是否能被3又能被5整除，打印提示信息
	var number1 int32 = 10
	var number2 int32 = 5
	if (number1+number2)%3 == 0 && (number1+number2)%5 == 0 {
		fmt.Println("可以被3和5同时整除")
	} else {
		fmt.Println("不能被3和5同时整除")
	}

	//	4.判断一个年份是否是闰年，闰年条件符合两个条件之一：（1.年份能被4整除但不能被100整除；2.能被400整除）
	var year int = 2020
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		fmt.Println("是闰年~")
	}
}
