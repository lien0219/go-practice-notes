package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	//	使用ioutil.ReadFile 一次性读取文件
	file := "D:\\桌面\\技术栈\\Golang\\go_project\\src\\go_code\\chapter14\\filedemo02\\file\\test.txt"
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("read file err=%v", err)
	}
	fmt.Printf("%v", string(content)) //[]byte
}
