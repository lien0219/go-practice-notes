package main

import "fmt"

func main() {

	//在go中，++和--只能独立使用

	//var i int = 8
	//var a int
	//a=i++ //错误，i++只能单独使用
	//a=i--  //错误，i--只能独立使用

	var i int = 1
	i++
	//++i   错误，go中没有前++
	i--
	//--i   错误，go中没有前--
	fmt.Println("i=", i)
}
