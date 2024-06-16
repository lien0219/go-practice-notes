package main

import "fmt"

// defer延时机制
func sum(n1 int, n2 int) int {
	//执行到defer时，暂时不执行，会将defer后面的语句压入独立的栈（defer栈）
	//当函数执行完毕后，再从defer栈中按照先入后出的方式出栈执行
	defer fmt.Println("n1=", n1)
	defer fmt.Println("n2=", n2)
	res := n1 + n2
	fmt.Println("res=", res)
	return res
}

func main() {
	result := sum(10, 20)
	fmt.Println("result=", result)
}
