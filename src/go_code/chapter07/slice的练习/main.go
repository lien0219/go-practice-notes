package main

import "fmt"

func test(slice []int) {
	slice[0] = 100 //这里修改会改变实参
}

func main() {
	var slice []int
	var arr [5]int = [...]int{1, 2, 3, 4, 5}
	slice = arr[:]
	var slice2 = slice
	slice2[0] = 10
	fmt.Println("slice2=", slice2) //10,2,3,4,5
	fmt.Println("slice=", slice)   //10,2,3,4,5
	fmt.Println("arr=", arr)       //10,2,3,4,5

	var slice4 = []int{1, 2, 3, 4}
	fmt.Println("slice4=", slice4) //1, 2, 3, 4
	test(slice4)
	fmt.Println("slice4=", slice4) //100,2,3,4
}
