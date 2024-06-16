package main

import "fmt"

func main() {
	//	切片的基本使用
	var intArr [5]int = [...]int{1, 22, 33, 66, 99}
	//	声明/定义切片
	//slice:=intArr[1:3]
	//1.slice就是切片名
	//2.intArr[1:3]表示slice引用到intArr这个数组
	//3.引用intArr数组的起始下标为1，最后的下标为3(但不包含3)
	slice := intArr[1:3]
	fmt.Println("intArr=", intArr)           //[1 22 33 66 99]
	fmt.Println("slice 的元素是=", slice)        //[22 33]
	fmt.Println("slice 的元素个数是=", len(slice)) //2
	fmt.Println("slice 的容量是=", cap(slice))   //容量可以动态变化  //4(一般是原数组的总长度-1)
}
