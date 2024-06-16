package main

import "fmt"

/*
go中的方法不仅仅是struct,还可以是int、float32等都可以有方法
*/
type integer int

func (i integer) print() {
	fmt.Println("i=", i)
}

// 编写一个方法，可以改变i的值
func (i *integer) change() {
	*i = *i + 1
}

type Student struct {
	Name string
	Age  int
}

// 给*Student实现方法String()
func (stu *Student) String() string {
	str := fmt.Sprintf("Name=[%v] Age=[%v]", stu.Name, stu.Age)
	return str
}

func main() {
	var i integer = 10
	i.print()
	i.change()
	fmt.Println("i=", i)

	stu := Student{
		Name: "tom",
		Age:  20,
	}
	fmt.Println(&stu)
}
