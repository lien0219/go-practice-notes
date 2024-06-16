package main

import "fmt"

// for循环
func main() {
	for i := 0; i <= 2; i++ {
		fmt.Println("for循环用法一", i)
	}

	j := 0
	for j <= 2 {
		fmt.Println("for循环用法二", j)
		j++
	}

	k := 0
	for {
		if k <= 2 {
			fmt.Println("for循环用法三", k)
		} else {
			break
		}
		k++
	}

	// 方式一	遍历字符串
	//注意：字符串中有中文的要转换成rune切片再循环
	var str string = "hello,world"
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c \n", str[i])
	}

	// 方式二	遍历字符串（for---range）
	str = "abc09"
	for index, val := range str {
		fmt.Printf("index=%d ,val=%c \n", index, val)
	}
}
