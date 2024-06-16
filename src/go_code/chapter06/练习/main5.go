package main

import "fmt"

/*
问题：
输出小写的a-z以及大写的Z-A
*/
/*
思路:
在Go中，字符（rune或byte，当用于ASCII字符时）是可以直接进行算术运算的
包括递增（++）和递减（--）。
在这个例子中，使用了rune类型（在Go中rune是int32的别名，用于表示Unicode字符），
但是当我们知道字符是ASCII字符时，使用byte（uint8的别名）也是可以的。
由于ASCII字符集在Go的rune和byte表示范围内，所以这里使用rune或byte都是可以的。
但是，为了代码的通用性和可读性，通常建议使用rune来表示字符，特别是当处理Unicode字符时。
*/

func main() {
	// 打印小写字母a-z
	for c := 'a'; c <= 'z'; c++ {
		fmt.Println(string(c))
	}
	fmt.Println()

	// 打印大写字母Z-A（需要递减）
	for c := 'Z'; c >= 'A'; c-- {
		fmt.Println(string(c))
	}
	fmt.Println()
}
