package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//在MethodUtils结构体编个方法，从键盘接收整数（1-9），打印对应的乘法表
/*
   MethodUtils 结构体：这只是一个用于封装方法的载体，不需要任何字段，我们只需要方法。
   从键盘接收输入可以使用bufio和os包中的Scanner来读取标准输入。
   打印乘法表：这涉及到遍历从1到用户输入的整数，并计算每个数与从1到用户输入数的乘积。
*/
type MethodUtils struct {
}

func (mu *MethodUtils) PrintMultiplicationTable() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("请输入一个整数（1-9）：")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("输入有错", err)
		return
	}
	// 去除换行符并转换为整数
	input = strings.TrimSpace(input)
	number, err := strconv.Atoi(input)
	if err != nil || number < 1 || number > 9 {
		fmt.Println("输入的不是有效整数（1-9）")
		return
	}
	// 打印乘法表
	for i := 1; i <= number; i++ {
		for j := 1; j <= number; j++ {
			fmt.Printf("%d*%d=%-2d\t", i, j, i*j)
		}
		fmt.Println()
	}
}

func main() {
	mu := MethodUtils{}
	mu.PrintMultiplicationTable()
}
