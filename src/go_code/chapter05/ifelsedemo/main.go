package main

import "fmt"

// if...else条件语句演示(双分支)
func main() {
	var age int
	fmt.Println("请输入你的年龄！")
	fmt.Scanln(&age)

	if age > 18 {
		fmt.Println("你的年龄大于18岁符合要求")
	} else {
		fmt.Println("你的年龄太小了！")
	}
}
