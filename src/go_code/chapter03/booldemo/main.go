package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var b = false
	fmt.Println("b=", b)
	//bool类型占用存储空间是一个字节
	fmt.Println("b 的占用空间 =", unsafe.Sizeof(b))
}
