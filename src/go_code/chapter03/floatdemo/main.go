package main

import (
	"fmt"
	"unsafe"
)

// 整数类型使用
func main() {
	var i int = 10
	fmt.Println("i=", i)

	//int8的范围是-128~127
	//其他的int16,int32,int64,类推
	var j int8 = -127
	fmt.Println("j=", j)

	//	uint8的范围是（0-255）,其他的uint16,uint32,uint64类推
	var k uint16 = 255
	fmt.Println("k=", k)

	//	int（有符号）,uint(无符号),rune,byte的使用
	var a int = 8900
	fmt.Println("a=", a)
	var b uint = 1
	var c byte = 255
	fmt.Println("b=", b, "c=", c)

	//	整数的使用细节
	var n1 = 100 //n1是什么类型？
	//	使用fmt.Printf()可以用于格式化输出
	fmt.Printf("n1的类型%T\n", n1)

	//	查看某个变量的占用字节大小和数据类型
	var n2 int64 = 10

	//unsafe.Sizeof(n2)  unsafe包的一个函数，可以返回n1变量占用的字节数
	fmt.Printf("n2 的 类型 %T  n2占用的字节数是 %d", n2, unsafe.Sizeof(n2))

	//	golang程序中整型变量在使用中，遵守保小不保大的原则
	//	尽量使用占用空间小的数据类型
	var age byte = 90
	fmt.Println("age=", age)
}
