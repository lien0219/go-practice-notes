package main

import "fmt"

// switch练习
func main() {
	//1.输出大小写
	//var char byte
	//fmt.Println("请输入一个字符")
	//fmt.Scanf("%c", &char)
	//switch char {
	//case 'a':
	//	fmt.Println("A")
	//case 'b':
	//	fmt.Println("B")
	//case 'c':
	//	fmt.Println("C")
	//case 'd':
	//	fmt.Println("D")
	//default:
	//	fmt.Println("other")
	//}

	//	2.根据输入成绩判断合格不合格
	//var score float64
	//fmt.Println("请输入成绩")
	//fmt.Scanln(&score)
	//switch {
	//case score >= 60:
	//	fmt.Println("合格")
	//case score < 60:
	//	fmt.Println("不合格")
	//}
	//switch int(score / 60) {
	//case 1:
	//	fmt.Println("合格")
	//case 0:
	//	fmt.Println("不合格")
	//default:
	//	fmt.Println("输入有误")
	//}

	//	根据输入判断季节
	var month byte
	fmt.Println("请输入月份")
	fmt.Scanln(&month)
	switch month {
	case 3, 4, 5:
		fmt.Println("spring")
	case 6, 7, 8:
		fmt.Println("summer")
	case 9, 10, 11:
		fmt.Println("autumn")
	case 12, 1, 2:
		fmt.Println("winter")
	default:
		fmt.Println("输入有误")
	}
}
