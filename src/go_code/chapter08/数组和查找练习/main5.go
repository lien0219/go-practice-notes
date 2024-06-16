package main

import "fmt"

/*
保存1，3，5，7，9五个奇数到数组，并倒序打印
*/
func main() {
	oddNumbers := []int{1, 3, 5, 7, 9}
	//方法一 直接使用for循环和len函数来倒序打印切片
	for i := len(oddNumbers) - 1; i >= 0; i-- {
		fmt.Print(oddNumbers[i], " ")
	}
	fmt.Println()

	//方法二 append和make创建一个新的切片，并复制原始切片的元素到新的切片中（倒序）
	reversedSlice := make([]int, len(oddNumbers))
	for i, num := range oddNumbers {
		reversedSlice[len(oddNumbers)-1-i] = num
	}
	fmt.Println("reversedSlice:", reversedSlice)
}
