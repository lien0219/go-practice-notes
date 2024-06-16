package main

import (
	"fmt"
	"strconv"
	_ "unsafe"
)

// 基本数据类型转string使用
func main() {
	var num1 int = 99
	var num2 float64 = 23.456
	var b bool = true
	var myChar byte = 'h'
	var str string //空string

	//	方法1.fmt.Sprintf转换成string
	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str type %T str=%v\n", str, str)

	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str type %T str=%v\n", str, str)

	str = fmt.Sprintf("%t", b)
	fmt.Printf("str type %T str=%v\n", str, str)

	str = fmt.Sprintf("%c", myChar)
	fmt.Printf("str type %T str=%q\n", str, str)

	//	方法2.strconv函数转换string
	var num3 int = 99
	var num4 float64 = 23.456
	var b2 bool = true

	str = strconv.FormatInt(int64(num3), 10)
	fmt.Printf("str type %T str=%q\n", str, str)
	//说明：'f' 格式   10：表示小数位保留十位    64：表示这个小数是float64
	str = strconv.FormatFloat(num4, 'f', 10, 64)
	fmt.Printf("str type %T str=%q\n", str, str)
	str = strconv.FormatBool(b2)
	fmt.Printf("str type %T str=%q\n", str, str)

	//	方法3.strconv包的Itoa函数
	var num5 int = 456
	str = strconv.Itoa(num5)
	fmt.Printf("str type %T str=%q\n", str, str)
	var num6 int64 = 4567
	str = strconv.Itoa(int(num6))
	fmt.Printf("str type %T str=%q\n", str, str)
}
