package main

import (
	"fmt"
	"time"
)

// 向intChan放1-8000个数
func putNum(intChan chan int) {
	for i := 1; i <= 80000; i++ {
		intChan <- i
	}
	close(intChan)
}

// 从intChan取出数据，判断是否为素数，是就放入primeChan
func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	//var num int
	var flag bool
	for {
		//time.Sleep(time.Millisecond * 10)
		num, ok := <-intChan
		if !ok {
			break
		}
		flag = true //假设是素数
		//	判断num是不是素数
		for i := 2; i < num; i++ {
			if num%i == 0 { //num不是素数
				flag = false
				break
			}
		}
		if flag {
			primeChan <- num
		}
	}
	fmt.Println("有一个primeNum协程因为取不到数据，退出")
	exitChan <- true
}

func main() {

	intChan := make(chan int, 1000)
	primeChan := make(chan int, 2000) //结果
	//	标识退出管道
	exitChan := make(chan bool, 4)

	start := time.Now().Unix()

	go putNum(intChan)
	for i := 0; i < 4; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}

		end := time.Now().Unix()
		fmt.Println("使用协程耗时为：", end-start)

		close(primeChan)
	}()

	for {
		res, ok := <-primeChan
		if !ok {
			break
		}
		fmt.Printf("素数为：%d\n", res)
	}
	fmt.Println("main线程退出")
}
