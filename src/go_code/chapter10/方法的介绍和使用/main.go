package main

import "fmt"

//go中的方法是作用在指定的数据类型上的（和指定的数据类型绑定）
//因此自定义类型都可以有方法，而不仅仅是struct

type Person struct {
	Name string
}

// 1.给Person结构体添加一个speak 方法，输出 xxx是好人
func (p Person) speak() {
	fmt.Println(p.Name, "是一个好人")
}

// 2.给Person结构体添加一个计算方法，可以计算1-1000的结果
func (p Person) jisuan() {
	res := 0
	for i := 0; i <= 1000; i++ {
		res += i
	}
	fmt.Println(p.Name, "计算的结果是=", res)
}

// 3.给Person结构体添加一个计算方法，该方法可以接收一个参数n,可以计算1+...+n的结果
func (p Person) jisuan2(n int) {
	res := 0
	for i := 0; i <= n; i++ {
		res += i
	}
	fmt.Println(p.Name, "计算的结果是=", res)
}

// 4.给Person结构体添加一个getSum方法，该方法可以计算两个数的和,并返回结果
func (p Person) getSum(n1 int, n2 int) int {
	return n1 + n2
}

// 给Person类型绑定一个方法 p表示哪个变量调用，这个p就是他的副本
func (p Person) test() {
	fmt.Println("test()", p.Name)
}

func main() {
	var p Person
	p.Name = "tom"
	p.test() //调用方法

	p.speak()
	p.jisuan()
	p.jisuan2(10)
	res := p.getSum(10, 20)
	fmt.Println("res=", res)
}
