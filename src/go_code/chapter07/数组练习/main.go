package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 练习
func main() {
	/*
		创建一个byte类型的26个元素的数组，分别放置’A‘-’Z‘
		使用for循环访问所有元素并打印出来，提示：字符数据运算'A'+1->'B'
	*/
	var myCharts [26]byte
	for i := 0; i < 26; i++ {
		myCharts[i] = 'A' + byte(i)
	}
	for i := 0; i < 26; i++ {
		fmt.Printf("%c ", myCharts[i])
	}
	fmt.Println()
	//	求出数组中的最大值并得到对应下标
	var intArr [5]int = [...]int{1, -1, 9, 90, 11}
	maxVal := intArr[0]
	maxValIndex := 0

	for i := 1; i < len(intArr); i++ {
		if maxVal < intArr[i] {
			maxVal = intArr[i]
			maxValIndex = i
		}
	}
	fmt.Printf("maxVal=%v maxValIndex=%v\n", maxVal, maxValIndex)
	fmt.Println()

	//	求出一个数组的和和平均值   for-range
	var intArr2 [5]int = [...]int{1, -1, 9, 90, 12}
	sum := 0
	for _, v := range intArr2 {
		sum += v
	}
	//平均值保留小数
	fmt.Printf("sum=%v 平均值=%v\n", sum, float64(sum)/float64(len(intArr2)))

	//	随机生成五个数，并反转打印
	var intArr3 [5]int
	lens := len(intArr3)
	//防止随机数每次创建都一样用seed种子初始化
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < lens; i++ {
		intArr3[i] = rand.Intn(100)
	}
	fmt.Println("交换前=", intArr3)
	//交换变量
	temp := 0
	for i := 0; i < lens/2; i++ {
		temp = intArr3[lens-1-i]
		intArr3[lens-1-i] = intArr3[i]
		intArr3[i] = temp
	}
	fmt.Println("交换后=", intArr3)
}
