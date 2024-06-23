package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//	打开文件
	//file 叫file对象 也可以叫 file指针 也可以叫 文件句柄
	file, err := os.Open("D:\\桌面\\技术栈\\Golang\\go_project\\src\\go_code\\chapter14\\filedemo02\\file\\test.txt")
	if err != nil {
		fmt.Println("open file err=", err)
	}

	fmt.Printf("file=%v \n", file) //file=&{0xc00009ca08}

	//	关闭文件
	defer file.Close()

	//	创建一个 *Reader ，是带缓冲的
	/*
		const (
			defaultBufSize=4096  //默认的缓冲区为4096
		)
	*/
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n') //读到换行结束
		if err == io.EOF {                  //io.EOF表示文件的末尾
			break
		}
		fmt.Print(str)
	}
	fmt.Println("文件读取结束...")
}
