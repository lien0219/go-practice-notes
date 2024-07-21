package main

import (
	"fmt"
	"reflect"
)

// 通过反射修改 num int的值，修改student的值
func reflect01(b interface{}) {
	rVal := reflect.ValueOf(b)
	//fmt.Printf("rVal kind: %v\n", rVal.Kind()) //ptr
	//Elem()获取指针指向的值
	rVal.Elem().SetInt(20) //修改num值
}
func main() {
	var num int = 10
	reflect01(&num)
	fmt.Println("num: ", num)
}
