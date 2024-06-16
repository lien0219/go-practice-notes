package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
	实现随机生成十个整数（1-100范围）保存到数组，并倒序打印以及求平均值、
	求最大值和最大值的下标、并查找里面是否有55
*/
/*
   打印原始切片
   倒序打印切片
   计算并打印平均值
   查找并打印最大值和最大值的下标
   查找并打印切片中是否包含数字55
*/
func main() {
	//	随机数
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 10)

	//	生成10个0-100范围的随机整数
	for i := range nums {
		nums[i] = rand.Intn(100) + 1
	}
	//	输出原始切片
	fmt.Println("原始切片：", nums)

	//	倒序打印切片
	for i := len(nums) - 1; i >= 0; i-- {
		fmt.Print(nums[i], " ")
	}
	fmt.Println()
	//	计算平均值
	sum := 0
	for _, num := range nums {
		sum += num
	}
	average := float64(sum) / float64(len(nums))
	fmt.Printf("平均值：%.2f\n", average)

	//	查找最大值和最大值下标
	maxNum := nums[0]
	maxIndex := 0
	for i, num := range nums {
		if num > maxNum {
			maxNum = num
			maxIndex = i
		}
	}
	fmt.Printf("最大值：%d, 下标：%d\n", maxNum, maxIndex)

	//	查找切片中是否有55
	contains55 := false
	for _, num := range nums {
		if num == 55 {
			contains55 = true
			break
		}
	}
	if contains55 {
		fmt.Println("切片中包含55")
	} else {
		fmt.Println("切片中不包含55")
	}
}
