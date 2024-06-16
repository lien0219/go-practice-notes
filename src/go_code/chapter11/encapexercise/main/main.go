package main

import (
	"fmt"
	"go_code/chapter11/encapexercise/model"
)

func main() {
	account := model.NewAccount("lien1111", "999999", 40)
	if account != nil {
		fmt.Println("创建成功=", *account)
	} else {
		fmt.Println("创建失败")
	}
}
