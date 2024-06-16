package main

import "fmt"

// 嵌套分支条件语句练习
func main() {
	//百米运动会，如果八秒以为进入决赛
	//	否则提前淘汰，根据性别提示进入男子组或者女子组
	//	输入性别
	//var second float64
	//fmt.Println("请输入秒数")
	//fmt.Scanln(&second)
	//
	//if second <= 8 {
	//	var gender string
	//	fmt.Println("请输入性别")
	//	fmt.Scanln(&gender)
	//	if gender == "男" {
	//		fmt.Println("进入男子组决赛")
	//	} else {
	//		fmt.Println("进入女子组决赛")
	//	}
	//
	//} else {
	//	fmt.Println("淘汰")
	//}

	/*
			出票系统：根据淡季旺季的月份和年龄，打印票价

			4-10 旺季
				成人（18-60）：60
				儿童（<18） ：半价
				老人(>60)  :1/3
		    淡季：
				成人：40
				其他：20
	*/

	var month byte
	var age byte
	var price float64 = 60.0
	fmt.Println("请输入月份")
	fmt.Scanln(&month)
	fmt.Println("请输入游客年龄")
	fmt.Scanln(&age)

	if month >= 4 && month <= 10 {
		if age > 60 {
			fmt.Printf("票价:%v", price/3)
		} else if age >= 18 {
			fmt.Printf("票价:%v", price)
		} else {
			fmt.Printf("票价:%v", price/2)
		}
	} else {
		if age >= 18 && age <= 60 {
			fmt.Printf("票价:%v", price)
		} else {
			price = 20.0
			fmt.Printf("票价:%v", price)
		}
	}
}
