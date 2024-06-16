package main

import "fmt"

func test() {
	//var n1 int = 10
}

func test2(n1 int) {
	n1 = n1 + 10
	fmt.Println("test2() n1=", n1)
}

// 引用传递值
func swap(x *int, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}

/* 值传递值 */
func sw(x, y int) int {
	var temp int

	temp = x /* 保存 x 的值 */
	x = y    /* 将 y 值赋给 x */
	y = temp /* 将 temp 值赋给 y*/

	return temp
}
func main() {
	//n1 := 20
	//test2(n1)
	//fmt.Println("mian() n1=", n1)

	/* 定义局部变量 */
	var a int = 100
	var b int = 200

	fmt.Printf("交换前，a 的值 : %d\n", a)
	fmt.Printf("交换前，b 的值 : %d\n", b)

	/* 调用 swap() 函数
	 * &a 指向 a 指针，a 变量的地址
	 * &b 指向 b 指针，b 变量的地址
	 */
	swap(&a, &b)

	fmt.Printf("交换后，a 的值 : %d\n", a)
	fmt.Printf("交换后，b 的值 : %d\n", b)

	/* 值传递值 */
	var c int = 10
	var d int = 20

	fmt.Printf("交换前 a 的值为 : %d\n", c)
	fmt.Printf("交换前 b 的值为 : %d\n", d)

	/* 通过调用函数来交换值 */
	sw(c, d)

	fmt.Printf("交换后 a 的值 : %d\n", c)
	fmt.Printf("交换后 b 的值 : %d\n", d)
}
