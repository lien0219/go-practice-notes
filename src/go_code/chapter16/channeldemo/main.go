package main

import "fmt"

func main() {
	//	创建一个可以存放3个int类型的管道
	var intChan chan int
	intChan = make(chan int, 3)
	fmt.Printf("intChan的值：%v intChan本身地址：%v \n", intChan, &intChan)

	//	向管道写入数据
	intChan <- 10
	num := 211
	intChan <- num
	intChan <- 50

	//	管道的长度和容量
	fmt.Printf("channel len:%v cap:%v \n", len(intChan), cap(intChan))

	//	从管道中读取数据
	var num2 int
	num2 = <-intChan
	fmt.Println("num2:", num2)

	//	管道的长度和容量
	fmt.Printf("channel len:%v cap:%v \n", len(intChan), cap(intChan))

	num3 := <-intChan
	num4 := <-intChan
	num5 := <-intChan //不能超过管道里数据的数量，否则报错
	fmt.Println("num3:", num3, " num4:", num4, " num5:", num5)
}
