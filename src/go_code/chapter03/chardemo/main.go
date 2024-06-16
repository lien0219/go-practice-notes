package main

import (
	"fmt"
)

func main() {

	var c1 byte = 'a'
	var c2 byte = '0'

	//	直接输出byte的值就是输出对应的字符的码值
	fmt.Println("c1:", c1)
	fmt.Println("c2:", c2)

	//	如果想输出对应的字符，需要用格式化输出
	fmt.Printf("c1=%c c2=%c\n", c1, c2)

	//var c3 byte = '测'   //溢出
	var c3 int = '测'
	fmt.Printf("c3=%c c3对应的码值=%d\n", c3, c3)

	//	在go中，字符的本质是一个整数，直接输出时，是该字符对应的UTF-8编码的码值
	//可以直接给某个变量赋值一个数字，然后按照格式化输出%c，会输出该数字对应的unicode字符
	var c4 int = 22269 //22269 ->国
	fmt.Printf("c4=%c\n", c4)

	//	字符类型是可以进行运算的，相当于一个整数，运算时是按照码值运行
	var n1 = 10 + 'a' //10+97=107
	fmt.Println("n1=", n1)

	//注：go语言的编码都统一成了utf-8了
}
