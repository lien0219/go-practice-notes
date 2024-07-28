package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "172.23.144.1:8888")
	defer conn.Close()
	if err != nil {
		fmt.Println("client dial err:", err)
		return
	}
	//fmt.Println("conn 成功：", conn)
	reader := bufio.NewReader(os.Stdin) //os.Stdin终端标准输入

	for {
		//	终端读取
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("readerString err:", err)
		}
		line = strings.Trim(line, " \r\n")
		if line == "exit" {
			fmt.Println("客户端退出！！！")
			break
		}
		//发送到服务器
		_, err = conn.Write([]byte(line + "\n"))
		if err != nil {
			fmt.Println("conn.Write err:", err)
		}
	}
	//fmt.Printf("客户端发送了%d字节的数据，并退出", n)
}
