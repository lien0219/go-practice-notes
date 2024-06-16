package main

import "fmt"

// for练习
func main() {
	//打印1-100之间所有的9的倍数的整数的个数以及总和
	var max uint64 = 100
	var count uint64 = 0
	var sum uint64 = 0

	var i uint64 = 1
	for ; i <= max; i++ {
		if i%9 == 0 {
			count++
			sum += i
		}
	}
	fmt.Printf("count=%v sum=%v\n", count, sum)

	fmt.Println("---------------------------")
	//	表达式输出，3是可变的
	var n int = 3
	for i := 0; i <= n; i++ {
		fmt.Printf("%v + %v = %v\n", i, n-i, n)
	}
}
