package main

import "fmt"

func main() {
	//第二种方式	演示切片的使用  make
	var slice []float64 = make([]float64, 5, 10)
	slice[1] = 10
	slice[2] = 20
	//	对于切片，必须make使用
	fmt.Println(slice)
	fmt.Println("slice的size=", len(slice))
	fmt.Println("slice的cap=", cap(slice))

	//	第三种方式  直接指定具体数组，使用原理类似make的方式
	var strSlice []string = []string{"tom", "jack", "mary"}
	fmt.Println("strSlice=", strSlice)
	fmt.Println("strSlice的size=", len(strSlice))
	fmt.Println("strSlice的cap=", cap(strSlice))

}
