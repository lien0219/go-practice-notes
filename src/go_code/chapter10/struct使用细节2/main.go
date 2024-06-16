package main

import (
	"encoding/json"
	"fmt"
)

type A struct {
	Num int
}

type B struct {
	Num int
}

type Monster struct {
	Name  string `json:"name"` //`json:"name"`就是struct tag 反射将Name转成小写name
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

func main() {
	var a A
	var b B
	a = A(b) //强制转换，要求必须结构体的字段要完全一样否则不行：Num int和Num int完全一样
	fmt.Println(a, b)

	//	1.Monster变量
	monster := Monster{"牛头", 500, "飞叉"}

	//	2.将monster变量序列化为json格式字符串
	jsonStr, err := json.Marshal(monster)
	if err != nil {
		fmt.Println("json处理错误", err)
	}
	// 3.jsonStr需要转为字符串
	fmt.Println("jsonStr:", string(jsonStr))
}
