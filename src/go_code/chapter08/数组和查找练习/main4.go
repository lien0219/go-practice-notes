package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
定义一个四行四列的二维数组，逐个从键盘输入值，然后将第一行和第四行的数据进行交换，将第二行和第三行数据进行交换
*/
func main() {
	//四行四列的二维数组
	var matrix [4][4]int

	// 从键盘逐个输入值
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			fmt.Printf("请输入第%d行第%d列的值: ", i+1, j+1)
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			num, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("输入错误，请输入一个整数")
				continue
			}
			matrix[i][j] = num
		}
	}
	//	原始二维数组
	fmt.Println("原始二维数组")
	for _, row := range matrix {
		fmt.Println(row)
	}

	//	交换第一行和第四行
	temp := matrix[0]
	matrix[0] = matrix[3]
	matrix[3] = temp

	// 交换第二行和第三行
	temp = matrix[1]
	matrix[1] = matrix[2]
	matrix[2] = temp

	fmt.Println("交换后的二维数组：")
	for _, row := range matrix {
		fmt.Println(row)
	}
}
