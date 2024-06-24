package main

import (
	"bufio"
	"fmt"
	"os"
)

// 1.打开一个存在的文件，将原来的内容覆盖成新的内容10句"你好，世界！"
func main() {
	/*
		const (
		    O_RDONLY int = syscall.O_RDONLY // 只读模式打开文件
		    O_WRONLY int = syscall.O_WRONLY // 只写模式打开文件
		    O_RDWR   int = syscall.O_RDWR   // 读写模式打开文件
		    O_APPEND int = syscall.O_APPEND // 写操作时将数据附加到文件尾部
		    O_CREATE int = syscall.O_CREAT  // 如果不存在将创建一个新文件
		    O_EXCL   int = syscall.O_EXCL   // 和O_CREATE配合使用，文件必须不存在
		    O_SYNC   int = syscall.O_SYNC   // 打开文件用于同步I/O
		    O_TRUNC  int = syscall.O_TRUNC  // 如果可能，打开时清空文件
		)
	*/

	//	打开文件
	filePath := "D:\\桌面\\技术栈\\Golang\\go_project\\src\\go_code\\chapter14\\filedemo05\\file\\writeTest.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0666) //0666 windows无作用
	if err != nil {
		fmt.Printf("open file err=%v \n", err)
		return
	}

	//	及时关闭防止内存泄露
	defer file.Close()

	//	写入5句 "hello,Garden"
	str := "你好，世界！ \r\n"
	//	写入时，使用带缓存的 *Writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}

	//	因为writer是带缓存的，因此在调用WriterString的时候
	//	其实内容是写入缓存的，所以要调用Flush方法，将带缓冲的数据
	//	真正写入到文件中，否则文件中没有数据
	writer.Flush()
}
