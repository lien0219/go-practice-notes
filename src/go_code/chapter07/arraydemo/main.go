package main

import "fmt"

// 演示for-range遍历数组
func main() {
	heroes := [...]string{"宋江", "狗头军师", "西门庆"}
	for i, v := range heroes {
		fmt.Printf("i=%v v=%v\n", i, v)
	}
}
