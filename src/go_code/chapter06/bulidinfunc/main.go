package main

import "fmt"

// new()函数解析
func main() {

	num1 := 100
	fmt.Printf("num1的类型%T，num1的值=%v，num1的地址%v\n", num1, num1, &num1)

	num2 := new(int)
	//num2的类型%T => *int
	//num2的值 =>地址  0xc00000a100
	//num2的地址 => 地址  0xc000054030
	//num2指针，指向的值 => 0

	//修改num2值
	*num2 = 100
	fmt.Printf("num2的类型%T，num2的值=%v，num2的地址%v\n num2指针，指向的值=%v", num2, num2, &num2, *num2)
}
