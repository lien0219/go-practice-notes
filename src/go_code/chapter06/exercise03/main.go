package main

import "fmt"

/*
有一堆桃子，猴子第一天吃了一半，并在多吃了一个，以后每天猴子都吃其中的一半，然后再多吃一个
当到第十天时，想再吃时（还没吃），发现只有一个桃子了。
问题：最初共有多少个桃子
*/
func peach(n int) int {
	if n > 10 || n < 1 {
		fmt.Println("输入的天数不对")
		return 0 //没有数量
	}
	if n == 10 {
		return 1
	} else {
		return (peach(n+1) + 1) * 2
	}
}

func main() {
	fmt.Println("第一天的桃子数量是：", peach(1))
}
