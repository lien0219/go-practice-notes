package main

import (
	"fmt"
	"math/rand"
	"time"
)

// break用于中断当前 for 循环或跳出 switch 语句
func main() {
	/*
		随机生成1-100的一个数，直到生成了99这个数，看看一共用了几次

	*/
	var count int = 0
	for {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(100) + 1
		count++
		if n == 99 {
			break
		}
	}
	fmt.Println("生成99 一共使用了 ", count)

	//	指定标签形式使用break

	//	不使用标记
	fmt.Println("----break----")
	for i := 1; i <= 3; i++ {
		fmt.Printf("i: %d\n", i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i2: %d\n", i2)
			break
		}
	}
	//	使用标记
	fmt.Println("---- break label ----")
re:
	for i := 1; i <= 3; i++ {
		fmt.Printf("i: %d\n", i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i2: %d\n", i2)
			break re
		}
	}
}
