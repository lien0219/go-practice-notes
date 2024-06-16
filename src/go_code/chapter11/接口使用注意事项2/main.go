package main

import "fmt"

type BInterface interface {
	test01()
}
type CInterface interface {
	test02()
}
type AInterface interface {
	BInterface
	CInterface
	test03()
}

// 如果需要实现AInterface,就需要将BInterface CInterface的方法都实现
type Stu struct {
}

func (stu Stu) test01() {}
func (stu Stu) test02() {}
func (stu Stu) test03() {}

type T interface {
}

func main() {
	//	8.一个接口（比如A接口）可以继承多个别的接口（比如B，C接口）这时如果要实现A接口，也必须将B，C接口的方法也全部实现
	var stu Stu
	var a AInterface = stu
	a.test01()
	//	9.interface类型默认是一个指针（引用类型），如果没有对interface初始化就是用，那么会输出nil
	//	10.空接口interface{}没有任何方法，所以所有类型都实现了空接口
	var t T = stu //ok
	fmt.Println(t)
	var t2 interface{} = stu //ok
	var num1 float64 = 8.8
	t2 = num1
	t = num1
	fmt.Println(t2, t)
}
