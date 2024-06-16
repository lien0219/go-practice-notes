package main

import "fmt"

/*
使用递归求出斐波那契数列
*/
func fbn(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return fbn(n-1) + fbn(n-2)
	}
}

func main() {
	res := fbn(3)
	fmt.Println("res=", res)
	fmt.Println("res=", fbn(4))
	fmt.Println("res=", fbn(5))
}
