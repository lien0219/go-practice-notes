package main

import (
	"fmt"
	"testing"
)

// 测试用例
func TestAddUpper(t *testing.T) {
	res := AddUpper(10)
	if res != 55 {
		//fmt.Printf("AddUpper(10) 执行错误，期望值：%v 实际值:%v \n", 55, res)
		t.Fatalf("AddUpper(10) 执行错误，期望值：%v 实际值:%v \n", 55, res)
	}
	t.Logf("AddUpper(10) 执行正确...")
}
func TestHello(t *testing.T) {
	fmt.Println("TestHello被调用...")
}
