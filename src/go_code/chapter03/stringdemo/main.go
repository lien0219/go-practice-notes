package main

import "fmt"

// string类型的使用
func main() {
	var address string = "浙江温州 110 hello world"
	fmt.Println(address)

	// 1.go中的字符串一旦赋值，字符串就不能修改了：go中的字符串是不可变的
	//var str = "hello"
	//str[0]="a" //go中的字符串是不可变的

	// 2.反引号表示字符串，可以以字符串原生形式输出，包括换行和特殊字符，可以防止攻击
	str2 := "abc\nabc"
	fmt.Println(str2)
	str3 := `
	str2 := "abc\nabc"
	fmt.Println(str2)
    `
	fmt.Println(str3)

	//	3.字符串的拼接
	var str = "hello" + "world"
	str += "haha!"

	fmt.Println(str)

	//	4.当拼接的操作很长时，可以分行写（ + 必须写后面）
	var str4 = "hello" + "world" + "hello" + "world" + "hello" + "world" +
		"hello" + "world" + "hello" + "world" +
		"hello" + "world"
	fmt.Println(str4)

	var a int           //0
	var b float32       //0
	var c float64       //0
	var isMarryied bool //false
	var name string     //''
	//fmt.Printf格式化输出     %v 表示按照变量的值输出
	fmt.Printf("a=%d,b=%f,c=%f,isMarryied=%v name=%v", a, b, c, isMarryied, name)
}
