package main

import "fmt"

func main() {
	var arr [5]int = [...]int{10, 20, 30, 40, 50}
	//slice := arr[1:4]
	slice := arr[0:len(arr)] //返回数组全部值 也可以简写为：arr[:]
	//使用for
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%v]=%v ", i, slice[i])
	}
	fmt.Println()
	//使用for-range
	for i, v := range slice {
		fmt.Printf("i=%v v=%v \n", i, v)
	}

	slice2 := slice[1:2] //切片可以再切
	slice2[0] = 100

	fmt.Println("slice2=", slice2)
	fmt.Println("slice=", slice)
	fmt.Println("arr=", arr)

	//	append内置函数，可以对切片动态添加
	var slice3 []int = []int{100, 200, 300}
	slice4 := append(slice3, 400, 500, 600)
	fmt.Println("slice3=", slice3)
	fmt.Println("slice4=", slice4)

	fmt.Println() //换行
	//切片的拷贝操作
	//切片使用copy内置函数完成拷贝
	var slice5 []int = []int{1, 2, 3, 4, 5}
	var slice6 = make([]int, 10)
	copy(slice6, slice5)           //参数的数据类型必须是切片
	fmt.Println("slice5=", slice5) //1，2，3，4，5
	fmt.Println("slice6=", slice6) //1，2，3，4，5，0，0，0，0，0  (独立空间，不会影响slice5)
}
