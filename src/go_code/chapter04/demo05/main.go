package main

import "fmt"

// 赋值运算符的使用
func main() {
	var i int

	i = 10
	fmt.Println("i=", i)

	a := 9
	b := 2
	fmt.Printf("交换前的情况是 a = %v , b = %v\n", a, b)

	t := a
	a = b
	b = t
	fmt.Printf("交换后的情况是 a = %v,b = %v\n", a, b)

	a += 19
	fmt.Println("a=", a)

	var c int
	c = a + 3
	fmt.Println("c=", c)

	var d int
	d = a
	d = 8 + 2*8
	fmt.Println("d=", d)
}
