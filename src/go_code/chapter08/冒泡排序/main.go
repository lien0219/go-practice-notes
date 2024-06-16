package main

import "fmt"

func BubbleSort(arr *[5]int) { //数组参数是值引用，如果修改原数组需要使用指针（引用地址）
	fmt.Println("排序前arr=", *arr)
	temp := 0
	for i := 0; i < len(*arr)-1; i++ {
		for j := 0; j < len(*arr)-1-i; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				//	交换
				temp = (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = temp
			}
		}
	}
	fmt.Println("排序后arr=", *arr)
	//测试
	////第一轮排序
	//for j := 0; j < 4; j++ {
	//	if (*arr)[j] > (*arr)[j+1] {
	//		//	交换
	//		temp = (*arr)[j]
	//		(*arr)[j] = (*arr)[j+1]
	//		(*arr)[j+1] = temp
	//	}
	//}
	//fmt.Println("第一次排序后arr=", *arr)
	////第二轮排序
	//for j := 0; j < 3; j++ {
	//	if (*arr)[j] > (*arr)[j+1] {
	//		//	交换
	//		temp = (*arr)[j]
	//		(*arr)[j] = (*arr)[j+1]
	//		(*arr)[j+1] = temp
	//	}
	//}
	//fmt.Println("第二次排序后arr=", *arr)
	////第三轮排序
	//for j := 0; j < 2; j++ {
	//	if (*arr)[j] > (*arr)[j+1] {
	//		//	交换
	//		temp = (*arr)[j]
	//		(*arr)[j] = (*arr)[j+1]
	//		(*arr)[j+1] = temp
	//	}
	//}
	//fmt.Println("第三次排序后arr=", *arr)
	////第四轮排序
	//for j := 0; j < 1; j++ {
	//	if (*arr)[j] > (*arr)[j+1] {
	//		//	交换
	//		temp = (*arr)[j]
	//		(*arr)[j] = (*arr)[j+1]
	//		(*arr)[j+1] = temp
	//	}
	//}
	//fmt.Println("第四次排序后arr=", *arr)

}

func main() {
	arr := [5]int{24, 69, 80, 57, 13}
	BubbleSort(&arr)
}
