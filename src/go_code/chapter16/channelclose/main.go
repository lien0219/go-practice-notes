package main

import "fmt"

func main() {
	intChan := make(chan int, 3)
	intChan <- 100
	intChan <- 200
	close(intChan) //close()

	//	关闭后不能再写入，但是可以读取
	//intChan <- 300
	fmt.Println("okokok")
	n1 := <-intChan
	fmt.Println("n1=", n1)

	/*   管道遍历只有一个值
	for item := range Chnl {
	     // 语句..
	}
	*/
	//	遍历管道
	intChan2 := make(chan int, 100)
	for i := 0; i < 100; i++ {
		intChan2 <- i * 2
	}

	//	关闭管道再进行遍历
	close(intChan2)

	//	遍历
	for v := range intChan2 {
		fmt.Println("v=", v)
	}

}
