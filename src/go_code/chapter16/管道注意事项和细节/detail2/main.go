package main

import "fmt"

func main() {

	//	使用select可以解决从管道取数据的阻塞问题

	//	1.定义一个管道10个数据int
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	//	2.定义一个管道5个数据string
	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "hello" + fmt.Sprintf("%d", i)
	}

	//	传统的遍历管道，如果不关闭会阻塞导致deadlock
	//	使用select方式解决
	for {
		select {
		//如果intChan一直没关闭，不会一直阻塞而deadlock
		//会自动到下一个case匹配
		case v := <-intChan:
			fmt.Printf("从intChan读取的数据%d\n", v)
		case v := <-stringChan:
			fmt.Printf("从stringChan读取的数据%s\n", v)
		default:
			fmt.Printf("都取不到。。。")
			return
		}
	}
}
