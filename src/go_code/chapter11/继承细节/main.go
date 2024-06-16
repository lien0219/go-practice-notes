package main

import "fmt"

type A struct {
	Name string
	age  int
}

func (a *A) SayOk() {
	fmt.Println("A SayOk", a.Name)
}

func (a *A) hello() {
	fmt.Println("A Hello", a.Name)
}

type B struct {
	A
	Name string
}

func (b *B) SayOk() {
	fmt.Println("B SayOk", b.Name)
}

func main() {
	//var b B
	//b.A.Name = "tom"
	//b.A.age = 19
	//b.A.SayOk()
	//b.A.hello()
	//
	////	上面写法简化
	//b.Name = "smith"
	//b.age = 20
	//b.SayOk()
	//b.hello()

	var b B
	b.Name = "jack" //ok
	b.A.Name = "scott"
	b.age = 100 //0k
	b.SayOk()   //B SayOk jack
	b.hello()
}
