package main

import "fmt"

// if条件语句演示(单分支)
func main() {

	//var age int
	//fmt.Println("请输入你的年龄：")
	////获取控制台输入的内容
	//fmt.Scanln(&age)
	//
	//if age > 18 {
	//	fmt.Println("你的年龄大于18岁，不符合条件！")
	//}

	//go中支持在if中直接定义变量(变量作用域仅限于当前if分支)
	if age := 20; age > 18 {
		fmt.Println("你的年龄大于18岁")
	}
}
