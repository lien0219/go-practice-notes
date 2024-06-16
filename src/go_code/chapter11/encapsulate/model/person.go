package model

import "fmt"

type person struct {
	name string
	age  int
	sal  float64
}

// 写个工厂模式的函数，相当于构造函数
func NewPerson(name string) *person {
	return &person{
		name: name,
	}
}

// 为了访问age和sal编写一个对SetXxxx的方法和GetXxx的方法
func (p *person) SetAge(age int) {
	if age > 0 && age < 150 {
		p.age = age
	} else {
		fmt.Println("年龄范围不对")
	}
}

func (p *person) SetSal(sal float64) {
	if sal >= 3000 && sal <= 30000 {
		p.sal = sal
	} else {
		fmt.Println("薪水范围不对")
	}
}
func (p *person) GetSal() float64 {
	return p.sal
}
