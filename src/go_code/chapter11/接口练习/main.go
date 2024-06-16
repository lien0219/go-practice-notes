package main

import "fmt"

// 测试1
type AInterface interface {
	Test01()
	Test02()
}
type BInterface interface {
	Test01()
	Test03()
}
type CInterface interface {
	AInterface
	BInterface
}

// 测试2
type Usb interface {
	Say()
}
type Stu struct {
}

func (this *Stu) Say() {
	fmt.Println("Say(")
}

func main() {
	//	测试2
	var stu Stu = Stu{}
	//var u Usb = stu //错误
	var u Usb = &stu //正确
	u.Say()
	fmt.Println("here", u)
}
