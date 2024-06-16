package main

import "fmt"

// Monkey结构体
type Monkey struct {
	Name string
}

// 声明一个接口
type BirdAble interface {
	Flying()
}
type FishAble interface {
	Swimming()
}

func (this *Monkey) climbing() {
	fmt.Println(this.Name, "爬树")
}

// LittleMonkey结构体
type LittleMonkey struct {
	Monkey //继承
}

// 让LittleMonkeys实现BirdAble
func (this *LittleMonkey) Flying() {
	fmt.Println(this.Name, "通过学习会飞翔！！！")
}

// 让LittleMonkeys实现FishAble
func (this *LittleMonkey) Swimming() {
	fmt.Println(this.Name, "通过学习会游泳！！！")
}

func main() {
	//	LittleMonkey实例
	monkey := LittleMonkey{
		Monkey{
			Name: "悟空",
		},
	}

	monkey.climbing()
	monkey.Flying()
	monkey.Swimming()
}
