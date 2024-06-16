package main

import "fmt"

/*
1.编写一个结构体（MethodUtils）,编写一个方法，方法不需要参数
在方法中打印10*8的矩形，在main中调用使用
*/
type MethodUtils struct {
	//	字段
}

func (mu MethodUtils) Print() {
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 8; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

/*
2.编写一个方法，提供m和n两个参数，方法中打印一个m*n的矩形
*/
func (mu MethodUtils) Print2(m int, n int) {
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

/*
3.编写一个方法,算该矩形的面积（可以接收长len和宽width）
将其作为方法返回值，在main方法中调用该方法，接收返回的面积值并打印
*/
func (mu MethodUtils) area(len float64, width float64) float64 {
	return len * width
}

// 4.编写方法，判断一个数是奇数还是偶数
func (mu *MethodUtils) JudgeNum(num int) {
	if num%2 == 0 {
		fmt.Println(num, "是偶数")
	} else {
		fmt.Println(num, "是奇数")
	}
}

// 5.根据行列字符打印对应的行数和列数的字符，比如：行：3，列：2，字符*，则打印相应的结果
func (mu *MethodUtils) Print3(n int, m int, key string) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			fmt.Print(key)
		}
		fmt.Println()
	}
}

// 6.定义一个计算器结构体（Calcuator）,实现加减乘除四个功能
type Calcuator struct {
	Num1 float64
	Num2 float64
}

// 方式一
func (calcuator *Calcuator) getSum() float64 {
	return calcuator.Num1 + calcuator.Num2
}
func (calcuator *Calcuator) getSub() float64 {
	return calcuator.Num1 - calcuator.Num2
}
func (calcuator *Calcuator) getCheng() float64 {
	return calcuator.Num1 * calcuator.Num2
}
func (calcuator *Calcuator) getChu() float64 {
	return calcuator.Num1 / calcuator.Num2
}

// 方式二
func (calcuator *Calcuator) getRes(operator byte) float64 {
	res := 0.0
	switch operator {
	case '+':
		res = calcuator.Num1 + calcuator.Num2
	case '-':
		res = calcuator.Num1 - calcuator.Num2
	case '*':
		res = calcuator.Num1 * calcuator.Num2
	case '/':
		res = calcuator.Num1 / calcuator.Num2
	default:
		fmt.Println("运算符输入有误...")
	}
	return res
}

func main() {
	var mu MethodUtils
	mu.Print()
	fmt.Println()
	mu.Print2(5, 20)

	areaRes := mu.area(2.5, 8.7)
	fmt.Println()
	fmt.Println("面积为：", areaRes)
	fmt.Println()
	(&mu).JudgeNum(3)

	(&mu).Print3(7, 20, "@")

	//	测试计算器
	var calcuator Calcuator
	calcuator.Num1 = 12.2
	calcuator.Num2 = 3.8
	getSum := (&calcuator).getSum()
	fmt.Printf("加：sum=%v", fmt.Sprintf("%.2f", getSum))
	fmt.Println()
	getSub := (&calcuator).getSub()
	fmt.Printf("减：sub=%v", fmt.Sprintf("%.2f", getSub))
	fmt.Println()

	res := (&calcuator).getRes('+')
	fmt.Println("res=", res)
}
