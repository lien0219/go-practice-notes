package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 5.将文件里的图片拷贝到另一个文件

// 辅助函数，接收两个文件路径
func CopyFile(dstFileName string, srcFileName string) (written int64, err error) {
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Printf("open file err=%v \n", err)
	}
	defer srcFile.Close()
	//	获取reader
	reader := bufio.NewReader(srcFile)
	//	打开dstFileName
	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err=%v \n", err)
		return
	}
	//	获取Writer
	writer := bufio.NewWriter(dstFile)

	defer dstFile.Close()

	return io.Copy(writer, reader)
}

func main() {
	//srcFile存在  dstFile不存在
	srcFile := "D:\\桌面\\技术栈\\Golang\\go_project\\src\\go_code\\chapter14\\filedemo05\\file\\ab\\test.jpg"
	dstFile := "D:\\桌面\\技术栈\\Golang\\go_project\\src\\go_code\\chapter14\\filedemo05\\file\\cd\\test.jpg"
	_, err := CopyFile(dstFile, srcFile)
	if err == nil {
		fmt.Printf("拷贝完成 \n")
	} else {
		fmt.Printf("拷贝错误 err=%v \n", err)
	}
}
