package main

import "fmt"

/*
使用递归求出f(n)的值
*/
func f(n int) int {
	if n == 1 {
		return 3
	} else {
		return 2*f(n-1) + 1
	}
}

func main() {
	fmt.Println("res=", f(1))
	fmt.Println("res=", f(5))
}
