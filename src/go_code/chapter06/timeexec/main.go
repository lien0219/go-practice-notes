package main

import (
	"fmt"
	"strconv"
	"time"
)

// 数组练习：计算函数执行耗费时间
func test03() {

	str := ""
	for i := 0; i < 100000; i++ {
		str += "hello" + strconv.Itoa(i)
	}
}

func main() {
	start := time.Now().Unix()
	test03()
	end := time.Now().Unix()
	fmt.Printf("执行test03耗时为%v秒\n", end-start)
}
