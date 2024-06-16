package main

import "fmt"

//编写一个函数，可以接收一个数组，该数组有五个数，请找出最大的数和最小的数和对应的数组下标是多少

func findMinMax(arr [5]int) (int, int, int, int) {
	minVal, minIdx := arr[0], 0
	maxVal, maxIdx := arr[0], 0
	for i, v := range arr {
		if v < minVal {
			minVal = v
			minIdx = i
		}
		if v > maxVal {
			maxVal = v
			maxIdx = i
		}
	}
	return maxVal, maxIdx, minVal, minIdx
}
func main() {
	arr := [5]int{3, 7, 1, 9, 2}
	maxVal, maxIdx, minVal, minIdx := findMinMax(arr)
	fmt.Printf("最大值是 %d 下标是 %d\n", maxVal, maxIdx)
	fmt.Printf("最小值是 %d 下标是 %d\n", minVal, minIdx)
}
