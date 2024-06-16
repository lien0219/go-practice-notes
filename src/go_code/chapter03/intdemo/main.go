package main

import "fmt"

func main() {

	var price float32 = 89.12
	fmt.Println("price =", price)
	var num1 float32 = -0.00089
	var num2 float64 = -7809656.09
	fmt.Println("num1=", num1, "num2=", num2)

	//	尾数部分丢失，造成精度丢失(float64的精度比float32的准确)
	var num3 float32 = -123.0000901
	var num4 float64 = -123.0000901
	fmt.Println("num3=", num3, "num4=", num4)

	//	golang的浮点型默认是float64
	var num5 = 1.1
	fmt.Printf("num5的数据类型是 %T\n", num5)

	//	十进制数形式必须有小数点
	num6 := 5.12
	num7 := .123
	fmt.Println("num6=", num6, "num7=", num7)

	//	科学计数形式
	num8 := 5.1234e2
	fmt.Println("num8=", num8)
}
