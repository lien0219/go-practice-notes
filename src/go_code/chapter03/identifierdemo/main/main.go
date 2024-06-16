package main

import (
	"fmt"
	_ "fmt" //如果没有使用一个包，但是想去掉，前面加一个 _ 表示忽略
	//src下
	"go_code/chapter03/identifierdemo/model"
)

// 标识符的使用
func main() {
	//1.会认为num和Num是不同变量(严格区分大小写)
	var num int = 10
	var Num int = 20
	fmt.Printf("num=%v Num=%v\n", num, Num)

	//2.标识符不能包含空格
	//var ab c int =30

	//3._是空标识符，用于占位
	var _ int = 40 //error
	//fmt.Println(_)

	/*
		标识符说明

		hello  //ok
		hello12 //ok
		1hello //error  不能以数字开头
		h-b   //error   不能使用-
		x h   //error    不能含有空格
		h_4   //ok
		_ab    //ok
		int   //ok   但不推荐
		float32   //ok  但不推荐
		_     //error
		Abc    //ok
		if  //error   保留关键字不行
	*/

	//go mod init (共同根终端执行一下才可以导入自定义包)
	fmt.Println(model.HeroName)
}
