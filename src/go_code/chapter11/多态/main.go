package main

import "fmt"

// 定义一个接口
type Usb interface {
	//声明两个没有实现的方法
	Start()
	Stop()
}
type Phone struct {
	name string
}
type Camera struct {
	name string
}

// 让phone实现Usb接口方法
func (p Phone) Start() {
	fmt.Println("手机开始工作")
}
func (p Phone) Stop() {
	fmt.Println("手机停止工作")
}

// 让Camera实现Usb接口方法
func (c Camera) Start() {
	fmt.Println("相机开始工作")
}
func (c Camera) Stop() {
	fmt.Println("相机停止工作")
}

func main() {
	//	定义一个Usb接口数组，可以存放Phone和Camera的结构体变量
	//	体现出多态数组
	var useArr [3]Usb
	useArr[0] = Phone{"vivo"}
	useArr[1] = Phone{"小米"}
	useArr[2] = Camera{"尼康"}

	fmt.Println(useArr)
}
