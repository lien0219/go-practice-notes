package main

import "fmt"

func getSum(n1 int, n2 int) int {
	return n1 + n2
}

// 函数既然是一种数据类型，函数可以作为形参，并且调用
func myFun(funvar func(int, int) int, num1 int, num2 int) int {
	return funvar(num1, num2)
}

// 案例
type myFunType func(int, int) int //这时myFunType就是 func(int, int) int类型
func myFun2(funvar myFunType, num1 int, num2 int) int {
	return funvar(num1, num2)
}

// 支持对函数的返回值命名
func getSumAndSub(n1 int, n2 int) (sum int, sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return
}

// 案例演示：编写一个函数sum,可以求出 1到多个int的和
// 可变参数演示（如果一个函数的形参列表有可变参数，要放在最后）
func sum(n1 int, args ...int) int {
	sum := n1
	for i := 0; i < len(args); i++ {
		sum += args[i] //args[0]表示取出args切片的第一个元素，其他的以此类推
	}
	return sum
}

func main() {
	a := getSum
	fmt.Printf("a的类型%T,getSum类型是%T\n", a, getSum)

	res := a(10, 40)
	fmt.Println("res=", res)

	res2 := myFun(getSum, 50, 60)
	fmt.Println("res2=", res2)

	//	给int取别名，在go中myInt和int虽然都是int类型，但是go中会认为是两种类型
	type myInt int

	var num1 myInt
	var num2 int
	num1 = 40
	num2 = int(num1) //依然要显示转换
	fmt.Println("num1=", num1, "num2=", num2)

	res3 := myFun2(getSum, 500, 600)
	fmt.Println("res3=", res3)

	//返回值命名案例
	a1, b1 := getSumAndSub(1, 2)
	fmt.Printf("a=%v, b=%v\n", a1, b1)

	//	测试可变参数的使用
	res4 := sum(10, 0, -1, 90, 10)
	fmt.Println("res4=", res4)
}
