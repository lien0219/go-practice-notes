package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	//  方式1

	//	方式2
	p2 := Person{
		"mary", 20,
	}
	fmt.Println(p2)

	//	方式3  指针形式  new   var person *Person=new(Person)
	var p3 *Person = new(Person)
	//因为p3是一个指针，所以要用以下方式赋值
	//(*p3).Name也可以这样写p3.Name="smith"  底层会进行处理，会给p3加上取值运算
	(*p3).Name = "smith"
	p3.Name = "john"

	(*p3).Age = 30
	p3.Age = 100

	fmt.Println(*p3)

	//	方式4   var person *Person=&Person{}
	var person *Person = &Person{}
	(*person).Name = "scott"
	person.Name = "lisi"

	(*person).Age = 88
	person.Age = 10
	fmt.Println(*person)
}
