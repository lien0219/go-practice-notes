package main

import (
	"fmt"
	_ "fmt" //如果没有使用一个包，但是想去掉，前面加一个 _ 表示忽略
)

// 测试
func main() {

	//var n1 int32 = 12
	//var n2 int64
	//var n3 int8
	//
	//n2 = int64(n1) + 20
	//n3 = int8(n1) + 20
	//
	//fmt.Println("n2=", n2, "n3=", n3)

	//var n1 int32 = 12
	////var n3 int8
	//var n4 int8
	//
	//n4 = int8(n1) + 127 //编译通过，但是结果按照溢出处理
	//fmt.Println(n4)

	//	指针测试
	/* var num int = 9

	fmt.Printf("num address=%v\n", &num)

	var ptr *int
	ptr = &num
	*ptr = 10
	fmt.Println("num=\n", num)  */

	//	数组练习
	var a int = 300                                               //ok
	var b int = 400                                               //ok
	var ptrr *int = &a                                            //ok
	*ptrr = 100                                                   // 100
	ptrr = &b                                                     //ok 400
	*ptrr = 200                                                   //200
	fmt.Printf("a=%d,b=%d,*ptrr=%d,ptrr=%d\n", a, b, *ptrr, ptrr) //a=100,b=200,*ptr=200
}
