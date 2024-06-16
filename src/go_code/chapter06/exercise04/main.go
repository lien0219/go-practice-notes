package main

import "fmt"

// 编写函数交换n1和n2值进行交换(不用指针是值传递，不会改变实际值)
func swap(n1 *int, n2 *int) {
	t := *n1
	*n1 = *n2
	*n2 = t
}

func main() {
	a := 10
	b := 20
	swap(&a, &b)
	fmt.Printf("a=%v, b=%v", a, b)
}
