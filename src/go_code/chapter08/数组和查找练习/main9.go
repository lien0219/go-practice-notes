package main

import "fmt"

// 实现定义一个数组，求该数组中大于平均值的数的个数，和小于平均值的数的个数
func countGreaterThanAverageAndLessThanAverage(arr []int) (int, int) {
	//	数组长度
	length := len(arr)
	if length == 0 {
		return 0, 0
	}

	//	数组总和
	sum := 0
	for _, val := range arr {
		sum += val
	}

	//	计算平均值
	average := float64(sum) / float64(length)
	fmt.Println("平均值:", average)

	//大于和小于平均值的计数器
	countGreater := 0
	countLess := 0
	for _, val := range arr {
		if float64(val) > average {
			countGreater++
		} else if float64(val) < average {
			countLess++
		}
	}
	return countGreater, countLess
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	countGreater, countLess := countGreaterThanAverageAndLessThanAverage(arr)
	fmt.Printf("大于平均值的数的个数: %d\n", countGreater)
	fmt.Printf("小于平均值的数的个数: %d\n", countLess)
}
