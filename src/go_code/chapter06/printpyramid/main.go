package main

import "fmt"

// 打印金字塔函数
func printPyramid(totalLevel int) {
	for i := 1; i <= totalLevel; i++ {
		//空格 总层数减去当前层数
		for k := 1; k <= totalLevel-i; k++ {
			fmt.Print(" ")
		}
		//2n-1 一行多少个星星
		for j := 1; j <= 2*i-1; j++ {
			//第一行和最后一行全打印星星，一行打几个剩余的空格
			if j == 1 || j == 2*i-1 || i == totalLevel {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

// 打印九九乘法表
func printmulti(multi int) {
	for i := 1; i <= multi; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v * %v = %v \t", j, i, j*i)
		}
		fmt.Println()
	}
}
func main() {
	var n int
	fmt.Println("请输入打印金字塔的层数")
	fmt.Scanln(&n)
	printPyramid(n)

	var num int
	fmt.Println("请输入打印乘法表的层数")
	fmt.Scanln(&num)
	printmulti(num)
}
