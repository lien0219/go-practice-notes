package main

import (
	"fmt"
	"go_code/chapter11/encapsulate/model"
)

func main() {
	p := model.NewPerson("史密斯")
	p.SetAge(18)
	p.SetSal(5000)
	fmt.Println(p)
}
