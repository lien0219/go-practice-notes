package main

import "fmt"

/*
案例：
编写一个Student结构体，包含name\gender\age\id\score字段
结构体中声明一个say方法，返回string类型 ，方法返回信息中包含所有字段值
在main方法中创建Student结构体实例，并访问say方法，将调用结果输出
*/
type Student struct {
	name   string
	gender string
	age    int
	id     int
	score  float64
}

func (student *Student) say() string {
	infoStr := fmt.Sprintf("student的信息 name=[%v] gender=[%v] age=[%v] id=[%v] score=[%v]",
		student.name, student.gender, student.age, student.id, student.score)
	return infoStr
}

/*
案例：

	创建一个Box结构体，在其中声明三个字段表示立方体的长宽高，长宽高要从终端获取
	声明一个方法获取立方体的体积
	创建一个Box立方体变量，打印给定尺寸的立方体的体积
*/
type Box struct {
	len    float64
	width  float64
	height float64
}

func (box *Box) getVolumn() float64 {
	return box.len * box.width * box.height
}

/*
	案例3：

一个景区根据游人的年龄收取不同价格的门票，比如年龄大于18，收费20，其他情况门票免费
编写一个Visitor结构体，根据年龄段决定能够购买的门票价格并输出
*/
type Visitor struct {
	Name string
	Age  int
}

func (visitor *Visitor) showPrice() {
	if visitor.Age >= 90 || visitor.Age <= 8 {
		fmt.Println("禁止入内")
		return
	}
	if visitor.Age > 18 {
		fmt.Printf("游客的名字为 %v 年龄为 %v 收费20元\n", visitor.Name, visitor.Age)
	} else {
		fmt.Printf("游客的名字为 %v 年龄为 %v 免费\n", visitor.Name, visitor.Age)
	}
}

func main() {
	//案例1
	var stu = Student{
		name:   "tom",
		gender: "male",
		age:    18,
		id:     1000,
		score:  99.98,
	}
	res := (&stu).say()
	fmt.Println(res)

	//	案例2
	var box Box
	box.len = 1.1
	box.width = 2.0
	box.height = 3.0
	volumn := box.getVolumn()
	fmt.Printf("体积为%.2f", volumn)

	//	案例3
	var v Visitor
	for {
		fmt.Println("请输入你的名字")
		fmt.Scanln(&v.Name)
		if v.Name == "n" {
			fmt.Println("退出程序。。。")
			break
		}
		fmt.Println("请输入你的年龄")
		fmt.Scanln(&v.Age)
		v.showPrice()
	}
}
