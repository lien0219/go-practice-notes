package main

import "fmt"

// 基本数据类型转换
func main() {

	var i int32 = 100
	//	希望将 i   =>float
	var n1 float32 = float32(i)
	var n2 int8 = int8(i)

	fmt.Printf("i=%v n1=%v n2=%v \n", i, n1, n2)

	//	被转换的是变量存储的数据（即值），变量本身的数据类型并没有变化
	fmt.Printf("i type is %T \n", i)

	//	在转换的过程中，比如将int64 转成 int8 [-128 ---127],编译时不会报错
	//只是转换的结果是按照溢出处理的，和希望结果不一样
	var num1 int64 = 999999
	var num2 int8 = int8(num1)
	fmt.Println("num2=", num2)
}
