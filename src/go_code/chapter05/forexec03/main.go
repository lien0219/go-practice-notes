package main

import "fmt"

func main() {

	// 练习打印空心金字塔
	var numC int = 10
	for i := 1; i <= numC; i++ {
		//空格 总层数减去当前层数
		for k := 1; k <= numC-i; k++ {
			fmt.Print(" ")
		}
		//2n-1 一行多少个星星
		for j := 1; j <= 2*i-1; j++ {
			//第一行和最后一行全打印星星，一行打几个剩余的空格
			if j == 1 || j == 2*i-1 || i == numC {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

	//	九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v * %v = %v \t", j, i, j*i)
		}
		fmt.Println()
	}
}
