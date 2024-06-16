package main

import "fmt"

type Person struct {
	Name string
}

// 函数
// 对于普通函数，接收者为值类型时，不能将指针类型的数据直接传递，反之亦然
func test01(p Person) {
	fmt.Println(p.Name)
}
func test02(p *Person) {
	fmt.Println(p.Name)
}

// 对于方法（struct）
// 接收者为值类型时，可以直接用指针类型的变量调用方法，反过来同样也可以
func (p Person) test03() {
	p.Name = "John"
	fmt.Println("test03=", p.Name) //John
}
func (p *Person) test04() {
	p.Name = "Mary"
	fmt.Println("test04=", p.Name) //Mary
}

func main() {
	p := Person{"tom"}
	test01(p)
	test02(&p)

	p.test03()
	fmt.Println("main() p.Name", p.Name) //tom

	(&p).test03() //形式上是传入地址 本质还是值拷贝

	fmt.Println("main() p.Name", p.Name) //tom

	(&p).test04()
	fmt.Println("main() p.Name", p.Name) //Mary
}
