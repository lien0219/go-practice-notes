package main

import "fmt"

// 二分查找演示
func BinaryFind(arr *[6]int, leftIndex int, rightIndex int, findVal int) {

	if leftIndex > rightIndex {
		fmt.Println("找不到")
		return
	}
	//	中间下标
	middle := (leftIndex + rightIndex) / 2
	if (*arr)[middle] > findVal {
		//	查找的数应该在 leftIndex----middle-1
		BinaryFind(arr, leftIndex, middle-1, findVal)
	} else if (*arr)[middle] < findVal {
		//	查找的数应该在 middle+1--- rightIndex
		BinaryFind(arr, middle+1, rightIndex, findVal)
	} else {
		//	找到了
		fmt.Printf("找到了下标为%v\n", middle)
	}
}

func main() {
	arr := [6]int{1, 8, 10, 89, 1000, 1234}
	BinaryFind(&arr, 0, len(arr)-1, 1000)
}
