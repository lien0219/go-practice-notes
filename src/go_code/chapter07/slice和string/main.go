package main

import "fmt"

func main() {
	//	string底层是一个byte数组，因此string也可以进行切片处理
	str := "hellolien" //不能str[0]=10,字符串不可变的
	slice := str[6:]
	fmt.Println("slice=", slice)

	//	演示怎么修改字符串 可以处理英文和数字但是不能处理中文
	//  原因是 []byte 字节处理，而一个汉字是三个字节，因此会出现乱码
	//  解决方法是将 string 转成 []rune 即可，因为 []rune 是按照字符处理的，兼容汉字

	//arr1 := []byte(str)
	//arr1[0] = 'z'
	//str = string(arr1)
	//fmt.Println("str=", str)
	arr1 := []rune(str)
	arr1[0] = '李'
	str = string(arr1)
	fmt.Println("str=", str)
}
