package main

import (
	"fmt"
	"sync"
	"time"
)

/*
计算1-200的各个数的阶乘，并且把各个数的阶乘放入到map中
*/
var (
	myMap = make(map[int]int, 10)
	//	全局互斥锁
	lock sync.Mutex
)

func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	//	加锁
	lock.Lock()
	myMap[n] = res
	//	解锁
	lock.Unlock()
}
func main() {
	for i := 1; i <= 20; i++ {
		go test(i)
	}
	time.Sleep(time.Second * 10)
	lock.Lock()
	for i, v := range myMap {
		fmt.Printf("map[%d]=%d\n", i, v)
	}
	lock.Unlock()
}
