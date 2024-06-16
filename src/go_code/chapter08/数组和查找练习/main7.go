package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
实现随机生成10个整数（1-100之间），使用冒泡排序法进行排序，
然后使用二分查找法查找是否有90这个数字，并显示下标，如果没有则提示找不到该数
*/

/*
生成10个1到100之间的随机整数。
使用冒泡排序法对数组进行排序。
使用二分查找法查找90这个数字，并显示下标（如果存在）。
如果没找到，则提示找不到该数。
*/

// 生成一个包含10个1到100之间随机整数的切片
func GenerateRandomSlice(n int) []int {
	rand.Seed(time.Now().UnixNano()) //初始化随机数种子
	slice := make([]int, n)
	for i := range slice {
		slice[i] = rand.Intn(100) + 1 //1到100的随机数
	}
	return slice
}

// 冒泡排序法对切片进行排序
func BubbleSort(slice []int) {
	n := len(slice)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
}

// 二分查找法在已排序的切片中查找目标值
func BinarySearch(slice []int, target int) (int, bool) {
	low, high := 0, len(slice)-1
	for low <= high {
		mid := (low + high) / 2
		if slice[mid] == target {
			return mid, true
		} else if slice[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1, false
}
func main() {
	randomSlice := GenerateRandomSlice(10)
	fmt.Println("随机数数组:", randomSlice)
	BubbleSort(randomSlice)
	fmt.Println("冒泡排序后:", randomSlice)
	index, found := BinarySearch(randomSlice, 90)
	if found {
		fmt.Printf("找到了，下标是: %d\n", index)
	} else {
		fmt.Println("没有找到")
	}
}
