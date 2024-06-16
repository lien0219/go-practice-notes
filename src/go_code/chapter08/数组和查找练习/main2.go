package main

import "fmt"

/*
已知有个排序好（升序）的数组，要求插入一个元素，最后打印该数组，顺序依然是升序
*/

func InsertElement(slice []int, element int) []int {
	insertIndex := len(slice)
	for i, v := range slice {
		if v >= element {
			insertIndex = i
			break
		}
	}
	// 创建一个新的切片，容量为原切片长度加1
	newSlice := make([]int, len(slice)+1)
	// 复制原切片中需要保留的元素到新切片中
	copy(newSlice[:insertIndex], slice[:insertIndex])
	// 将新元素插入到正确的位置
	newSlice[insertIndex] = element
	// 如果需要，复制原切片中剩余的元素到新切片中
	copy(newSlice[insertIndex+1:], slice[insertIndex:])
	return newSlice
}
func main() {
	/*
		遍历数组以找到应该插入新元素的位置（即第一个大于或等于新元素的元素的位置）。
		将这个位置及其后面的所有元素向后移动一位，为新元素腾出空间。
		在找到的位置插入新元素。
		打印整个数组
	*/
	sortedSlice := []int{1, 3, 5, 6, 9}
	//	要插入的元素
	elementToInsert := 4
	sortedSlice = InsertElement(sortedSlice, elementToInsert)
	fmt.Println(sortedSlice) // 输出: [1 3 4 5 7 9]
}
