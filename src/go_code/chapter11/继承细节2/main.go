package main

import "fmt"

type A struct {
	Name string
	age  int
}

type B struct {
	Name  string
	Score float64
}

type C struct {
	A
	B
	//Name string
}

type D struct {
	a A //有名结构体
}

type Goods struct {
	Name  string
	Price float64
}
type Brand struct {
	Name    string
	Address string
}

type TV struct {
	*Goods
	*Brand
}
type Monster struct {
	Name string
	Age  int
}
type E struct {
	Monster
	int
}

func main() {
	var c C
	//如果c没有Name字段，而A、B有Name,这时必须指定匿名结构体名字区分
	c.A.Name = "tom"
	fmt.Println("c")

	var d D
	d.a.Name = "jack"

	//	嵌套匿名结构体后，也可以在创建结构体变量时，直接指定各个匿名结构体字段的值
	tv := TV{&Goods{"电视机001", 5000.99}, &Brand{"海尔", "山东"}}
	fmt.Println("tv:", *tv.Goods, *tv.Brand)
	tv2 := TV{
		&Goods{
			Price: 500.99,
			Name:  "电视机002",
		},
		&Brand{
			Name:    "夏普",
			Address: "河南",
		},
	}
	fmt.Println("tv2:", *tv2.Goods, *tv2.Brand)

	//	演示匿名字段是基本数据类型的使用
	var e E
	e.Name = "狐狸"
	e.Age = 300
	e.int = 20
	fmt.Println("e=", e)
}
