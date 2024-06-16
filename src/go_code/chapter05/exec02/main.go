package main

import (
	"fmt"
	"math"
)

// 多分支条件语句练习
func main() {
	/*
			小明参加考试：成绩为100分时，奖励宝马一辆，
			成绩为80-99时奖励一台iphone手机，成绩为60-80时，奖励ipad
			其他成绩则不奖励
		从键盘输入获取成绩
	*/
	var score int
	fmt.Println("请输入成绩！")
	fmt.Scanln(&score)
	if score == 100 {
		fmt.Println("奖励宝马一辆")
	} else if score > 80 && score <= 99 {
		fmt.Println("iphone手机")
	} else if score >= 60 && score <= 80 {
		fmt.Println("奖励ipad")
	} else {
		fmt.Println("没有奖励")
	}

	var a float64 = 3.0
	var b float64 = 100.0
	var c float64 = 6.0

	m := b*b - 4*a*c

	if m > 0 {
		x1 := (-b + math.Sqrt(m)) / 2 * a
		x2 := (-b - math.Sqrt(m)) / 2 * a
		fmt.Printf("x1=%v x2=%v \n", x1, x2)
	} else if m == 0 {
		x1 := (-b + math.Sqrt(m)) / 2 * a
		fmt.Printf("x1=%v", x1)
	} else {
		fmt.Println("无解。。。")
	}

	//	嫁娶
	var height int32
	var money float32
	var handsome bool

	fmt.Println("请输入身高（cm）")
	fmt.Scanln(&height)
	fmt.Println("请输入财富（千万）")
	fmt.Scanln(&money)
	fmt.Println("请输入是否帅（true/false）")
	fmt.Scanln(&handsome)

	if height > 180 && money > 1.0 && handsome {
		fmt.Println("一定嫁给他")
	} else if height > 180 || money > 1.0 || handsome {
		fmt.Println("考虑一下")
	} else {
		fmt.Println("滚吧")
	}
}
