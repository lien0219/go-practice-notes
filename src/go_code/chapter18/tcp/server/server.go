package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	//	接收客户端数据
	defer conn.Close()

	for {
		buf := make([]byte, 1024)
		//fmt.Printf("等待客户端%s发送信息!!!\n", conn.RemoteAddr().String())
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("客户端退出 err:%v", err)
			return
		}
		fmt.Print(string(buf[:n]))
	}
}

func main() {

	fmt.Println("服务器开始监听。。。。。。")
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen err:", err)
		return
	}

	defer listen.Close() //延迟关闭

	for {
		fmt.Println("等待客户端连接......")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept() err:", err)
		} else {
			fmt.Printf("Accept() success conn:%v 客户端ip:%v\n", conn, conn.RemoteAddr().String())
		}
		go process(conn)
	}

	//fmt.Printf("listen success=%v \n", listen)
}
