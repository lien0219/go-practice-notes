package main

import "fmt"

type Cat struct {
	Name  string
	Age   int
	Color string
	Hobby string
}

func main() {
	/*
		养两只猫：一只名字叫小白，今年三岁，白色；还有一只叫小花，
		今年100岁，花色。请编写一个程序，当用户输入小猫的名字时，显示该猫的名字
		年龄，颜色，如果用户输入的小猫名字有误则显示没有这只猫
	*/

	//使用变量处理（麻烦不建议）
	//var catName string = "小白"
	//var catAge int = 3
	//var catColor string = "白色"
	//
	//var cat2Name string = "小花"
	//var cat2Age int = 100
	//var cat2Color string = "花色"

	//	使用数组处理（麻烦不建议）
	//var catNames [2]string = [...]string{"小白", "小花"}
	//var catAges [2]int = [...]int{3, 100}
	//var catColors [2]string = [...]string{"白色", "花色"}、

	//	使用结构体解决（struct）
	//创建cat变量
	var cat1 Cat
	cat1.Name = "小白"
	cat1.Age = 3
	cat1.Color = "白色"
	cat1.Hobby = "吃鱼"

	fmt.Println("cat1:")
	fmt.Println(cat1)

	fmt.Println("猫的信息如下：")
	fmt.Println("name=", cat1.Name)
	fmt.Println("age=", cat1.Age)
	fmt.Println("color=", cat1.Color)
	fmt.Println("hobby=", cat1.Hobby)
}
