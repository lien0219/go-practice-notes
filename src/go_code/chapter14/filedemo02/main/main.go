package main

import (
	"fmt"
	"os"
)

func main() {
	//	打开文件
	//file 叫file对象 也可以叫 file指针 也可以叫 文件句柄
	file, err := os.Open("D:\\桌面\\技术栈\\Golang\\go_project\\src\\go_code\\chapter14\\filedemo02\\file\\test.txt")
	if err != nil {
		fmt.Println("open file err=", err)
	}

	fmt.Printf("file=%v", file) //file=&{0xc00009ca08}

	//	关闭文件
	err = file.Close()
	if err != nil {
		fmt.Println("close file err=", err)
	}

}
