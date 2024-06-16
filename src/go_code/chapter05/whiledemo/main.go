package main

import "fmt"

func main() {
	// while方式循环
	var i int = 1
	for {
		if i > 3 {
			break
		}
		fmt.Println("hello,world", i)
		i++
	}
	fmt.Println("i=", i) //4

	//	do...while方式循环
	var j int = 1
	for {
		fmt.Println("哈哈哈", j)
		j++
		if j > 3 {
			break
		}
	}
	fmt.Println("j=", j) //4
}
