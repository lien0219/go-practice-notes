package main

import "fmt"

// fmt.Scanln和fmt.Scanf获取控制台用户输入信息
func main() {
	//1.从控制台接收用户信息，（姓名、年龄、薪水、是否通过考试）
	//	方式1   fmt.Scanln
	var name string
	var age byte
	var sal float32
	var isPass bool
	//fmt.Println("请输入姓名")
	////执行到此时会等待用户输入
	//fmt.Scanln(&name)
	//fmt.Println("请输入年龄")
	////执行到此时会等待用户输入
	//fmt.Scanln(&age)
	//fmt.Println("请输入薪水")
	////执行到此时会等待用户输入
	//fmt.Scanln(&sal)
	//fmt.Println("请输入是否通过考试")
	////执行到此时会等待用户输入
	//fmt.Scanln(&isPass)
	//
	//fmt.Printf("名字是 %v \n 年龄是 %v \n 薪水是 %v \n 是否通过考试 %v \n", name, age, sal, isPass)

	//方式2：fmt.Scanf,可以按照指定格式输入
	fmt.Println("请输入你的姓名、年龄、薪水、是否通过考试，使用空格隔开")
	fmt.Scanf("%s %d %f %t", &name, &age, &sal, &isPass) //具体含义文档 https://studygolang.com/pkgdoc
	fmt.Printf("名字是 %v \n 年龄是 %v \n 薪水是 %v \n 是否通过考试 %v \n", name, age, sal, isPass)
}
