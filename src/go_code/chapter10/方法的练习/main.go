package main

import "fmt"

//1.声明一个结构体
//2.声明一个方法area和Circle绑定，可以返回面积

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

// 提高效率，方法和结构体的指针类型绑定
func (c *Circle) area2() float64 {
	return 3.14 * (*c).radius * (*c).radius
}

func main() {
	//var c Circle
	//c.radius = 4.0
	//res := c.area()
	//fmt.Println("面积是：", res)

	var c Circle
	c.radius = 5.0
	res := (&c).area2()
	//编译器底层做了优化   (&c).area2()等价于c.area2()
	fmt.Println("面积是：", res)
}
