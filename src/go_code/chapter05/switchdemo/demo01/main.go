package main

import "fmt"

// switch案例
func main() {
	/*
			switch:
		        switch 语句用于基于不同条件执行不同动作，每一个 case 分支都是唯一的，从上至下逐一测试，直到匹配为止。
		        switch 语句执行的过程从上至下，直到找到匹配项，匹配项后面也不需要再加 break。
		        switch 默认情况下 case 最后自带 break 语句，匹配成功后就不会执行其他 case，如果我们需要执行后面的 case，可以使用 fallthrough 。

					语法
					Go 编程语言中 switch 语句的语法如下：
				switch var1 {
				    case val1:
				        ...
				    case val2:
				        ...
				    default:
				        ...
				}

			变量 var1 可以是任何类型，而 val1 和 val2 则可以是同类型的任意值。类型不被局限于常量或整数，但必须是相同的类型；或者最终结果为相同类型的表达式。
			您可以同时测试多个可能符合条件的值，使用逗号分割它们，例如：case val1, val2, val3。
	*/

	/*
		案例：	编写程序，该程序接收一个字符，例如a,b,c,d,e,f,g
			a表示星期一，b表示星期二。。。。根据用户输入显示相应信息
	*/

	//var key byte
	//fmt.Println("请输入一个字符a,b,c,d,e,f,g")
	//fmt.Scanf("%c", &key)
	//
	//switch key {
	//case 'a':
	//	fmt.Println("周一")
	//case 'b':
	//	fmt.Println("周二")
	//case 'c':
	//	fmt.Println("周三")
	//case 'd':
	//	fmt.Println("周四")
	//case 'e':
	//	fmt.Println("周五")
	//case 'f':
	//	fmt.Println("周六")
	//case 'g':
	//	fmt.Println("周日")
	//default:
	//	fmt.Println("输入错误。。。")
	//}

	var n1 int32 = 20
	var n2 int32 = 20
	switch n1 {
	case n2: //必须是同类型
		fmt.Println("ok")
	}

	var n3 int32 = 5
	var n4 int32 = 19
	switch n3 {
	case n4, 10, 5: //case后面可以有多个表达式用逗号隔开
		fmt.Println("ok111")
	}

	//	switch 后面可以不带表达式，类似于if else 分支
	var age int = 10
	switch {
	case age == 10:
		fmt.Println("age=10")
	case age == 20:
		fmt.Println("age=20")
	default:
		fmt.Println("匹配错误")
	}

	//数组练习
	//var grade string = "B"
	//var marks int
	//
	//fmt.Println("请输入考试分数")
	//fmt.Scanln(&marks)
	//switch marks {
	//case 90:
	//	grade = "A"
	//case 80:
	//	grade = "B"
	//case 50, 60, 70:
	//	grade = "C"
	//default:
	//	grade = "D"
	//}
	//switch {
	//case grade == "A":
	//	fmt.Printf("优秀!\n")
	//case grade == "B", grade == "C":
	//	fmt.Printf("良好\n")
	//case grade == "D":
	//	fmt.Printf("及格\n")
	//case grade == "F":
	//	fmt.Printf("不及格\n")
	//default:
	//	fmt.Printf("差\n")
	//}
	//fmt.Printf("你的等级是 %s\n", grade)

	//	switch 的穿透 fallthrough
	var num int = 10
	switch num {
	case 10:
		fmt.Println("10")
		fallthrough //默认只能穿透一层
	case 20:
		fmt.Println("20")
		fallthrough //默认只能穿透一层
	case 30:
		fmt.Println("30")
	default:
		fmt.Println("匹配错误")
	}
}
