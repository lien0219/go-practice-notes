package main

import "fmt"

// Go 语言的 goto 语句可以无条件地转移到过程中指定的行。
func main() {

	var n int = 5
	//	演示goto的使用
	fmt.Println("pk1")
	if n > 10 {
		goto label1
	}
	fmt.Println("pk2")
	fmt.Println("pk3")
	fmt.Println("pk4")
label1:
	fmt.Println("pk5")
	fmt.Println("pk6")
	fmt.Println("pk7")
}
