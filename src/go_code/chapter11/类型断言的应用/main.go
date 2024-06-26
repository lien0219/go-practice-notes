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
func (p Phone) Call() {
	fmt.Println("手机 在打电话。。。")
}

// 让Camera实现Usb接口方法
func (c Camera) Start() {
	fmt.Println("相机开始工作")
}
func (c Camera) Stop() {
	fmt.Println("相机停止工作")
}

type Computer struct {
}

func (computer Computer) Working(usb Usb) {
	usb.Start()
	//如果usb是指向Phone结构体变量，则还需要调用Call方法
	//类型断言
	if phone, ok := usb.(Phone); ok == true {
		phone.Call()
	}
	usb.Stop()
}

func main() {
	//	定义一个Usb接口数组，可以存放Phone和Camera的结构体变量
	//	体现出多态数组
	var useArr [3]Usb
	useArr[0] = Phone{"vivo"}
	useArr[1] = Phone{"小米"}
	useArr[2] = Camera{"尼康"}
	fmt.Println(useArr)

	//遍历useArr
	//Phone还有一个特有的方法Call(),请遍历Usb数组，如果是Phone变量
	//除了调用Usb接口声明的方法外，还需要调用Phone特有方法Call() =>类型断言
	var computer Computer
	for _, value := range useArr {
		computer.Working(value)
		fmt.Println()
	}

}
