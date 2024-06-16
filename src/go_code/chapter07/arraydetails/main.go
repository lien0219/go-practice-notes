package main

import "fmt"

func test01(arr [3]int) {
	arr[0] = 88
}
func test02(arr *[3]int) {
	(*arr)[0] = 88
}

// 数组的注意事项
func main() {
	//  1.数组是多个相同类型数据的组合，一个数组一旦声明/定义了，其长度是固定的，不能动态变化
	var arr01 [3]int
	arr01[0] = 1
	arr01[1] = 2
	arr01[2] = 3
	fmt.Println(arr01)

	//	2.var arr []int  不写长度[]为空，arr就是一个slice切片，后续

	//	3.数组中的元素可以是任何数据类型，包括值类型和引用类型，但是不能混用
	var arr02 [3]int
	//arr01[0] = 1.1 //报错，数组类型是int类型
	arr01[0] = 1
	arr01[1] = 2
	arr01[2] = 3
	fmt.Println(arr02)

	//	4.数组创建完成后，如果没有赋值，有默认值0
	//	4.1、数值（整数系列和浮点数系列）=》0
	//	4.2、字符串=》“”
	//	4.3、布尔类型=》false
	var arr03 [3]float32
	var arr04 [3]string
	var arr05 [3]bool
	fmt.Printf("arr03=%v arr04=%v arr05=%v\n", arr03, arr04, arr05)

	//	5.使用数组的步骤1.声明数组开辟空间2.给各个元素赋值3.使用数组

	//	6.数组的下标是从0开始的
	//var arr06 [3]string
	//var index int = 3
	//arr06[index] = "3" //超出界限，最大长度为2

	//	7.数组下标必须在指定范围内使用，否则报panic：数组越界，例如：var arr[5]int 则有效下标为0-4

	//	8.数组属值类型，在默认情况下是值传递，因此会进行值拷贝，数组间不会相互影响
	arr := [3]int{11, 22, 33}
	test01(arr)
	fmt.Println("arr=", arr)

	//	9.如果想在其他函数中修改原来的数组，可以使用引用传递（指针方式）
	arr07 := [3]int{1, 2, 3}
	test02(&arr07)
	fmt.Println("arr07=", arr07)

	//	10.长度是数组类型的一部分，在传递函数参数时需要考虑数组的长度
}
