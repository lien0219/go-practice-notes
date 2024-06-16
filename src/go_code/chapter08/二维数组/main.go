package main

import "fmt"

// 二维数组演示
func main() {
	//	定义二维数组
	var arr [4][6]int

	arr[1][2] = 1
	arr[2][1] = 2
	arr[2][3] = 3

	for i := 0; i < 4; i++ {
		for j := 0; j < 6; j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}

	//fmt.Println("arr=", arr)

	var arr2 [2][3]int
	arr2[1][1] = 10
	fmt.Println(arr2)

	var arr3 [2][3]int = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("arr3=", arr3)

	var arr4 [2][3]int = [...][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("arr4=", arr4)
}
