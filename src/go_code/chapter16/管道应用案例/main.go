package main

import (
	"fmt"
)

func writeData(intChan chan int) {
	for i := 1; i <= 50; i++ {
		//	数据放入管道z
		intChan <- i
		fmt.Println("writeData:", i)
		//方便直观的看协程效果
		//time.Sleep(time.Second)
	}
	close(intChan)
}
func readData(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		//方便直观的看协程效果
		//time.Sleep(time.Second)
		fmt.Printf("readData读取数据：%v\n", v)
	}
	//	数据读取完毕
	exitChan <- true
	close(exitChan)
}
func main() {
	//	创建两个管道
	intChan := make(chan int, 50)
	exitChan := make(chan bool, 1)

	//	两个协程
	go writeData(intChan)
	go readData(intChan, exitChan)

	//	休眠防止主线程完毕协程关闭(线程礼让)
	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}
}
