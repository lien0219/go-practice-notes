package main

import "fmt"

// continue:Go 语言的 continue 语句 有点像 break 语句。
// 但是 continue 不是跳出循环，而是跳过当前循环执行下一次循环语句。
func main() {

	//	continue
	for i := 1; i < 4; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				continue
			}
			fmt.Println("j=", j)
		}
	}
}
