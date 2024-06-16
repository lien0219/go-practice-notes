package main

import (
	"fmt"
	"go_code/chapter10/factory/model"
)

func main() {
	//	创建Student实例
	//var stu = model.Student{
	//	Name:  "tom",
	//	Score: 19.8,
	//}

	//当student结构体是首字母小写，可以通过工厂模式解决
	var stu = model.NewStudent("tom~", 88.8)
	fmt.Println(*stu)
	//fmt.Println("name=", (*stu).Name, "score=", (*stu).Score)
	fmt.Println("name=", (*stu).Name, "score=", (*stu).GetScore())
}
