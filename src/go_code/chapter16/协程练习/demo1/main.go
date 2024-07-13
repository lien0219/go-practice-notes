package main

import (
	"fmt"
	"sync"
)

/*
启动一个协程，将1-2000放入channel,(numChan)
启动八个协程，从numChan取出数并计算相加的值放入resChan
八个协程同时完成工作再遍历resChan,显示结果如：res[1]=1...res[10]=55
需要考虑resChan chan int是否合适
*/

func main() {
	numChan := make(chan int, 2000)
	resChan := make(chan int, 8)

	var wg sync.WaitGroup
	//一个协程发送数字
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 2000; i++ {
			numChan <- i
			fmt.Println("发送数字：", i)
		}
		close(numChan)
	}()

	// 8个协程处理数字
	for i := 1; i <= 8; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			sum := 0
			for num := range numChan {
				sum += num
			}
			resChan <- sum
		}(i)
	}

	//	等待所有的协程完成
	wg.Wait()

	//演示接收
	results := make(map[int]int)
	for i := 0; i < 8; i++ {
		sum := <-resChan
		// 假设每个goroutine的id就是结果数组的索引
		results[i+1] = sum //所有数的总和
	}

	//	打印结果
	for k, v := range results {
		fmt.Printf("res[%d]=%d\n", k, v)
	}
}
