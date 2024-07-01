package main

import "testing"

func TstSub(t *testing.T) {
	res := getSub(10, 3)
	if res != 7 {
		//fmt.Printf("AddUpper(10) 执行错误，期望值：%v 实际值:%v \n", 55, res)
		t.Fatalf("getSub(10, 3) 执行错误，期望值：%v 实际值:%v \n", 7, res)
	}
	t.Logf("getSub(10, 3) 执行正确...")
}
