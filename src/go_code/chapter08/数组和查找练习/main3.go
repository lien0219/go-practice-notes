package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
	定义一个三行四列的二维数组，逐个从键盘输入值，编写程序将四周的数据清0
*/

func main() {
	//	初始化三行四列二维切片
	matrix := make([][]int, 3)
	for i := range matrix {
		matrix[i] = make([]int, 4)
	}

	//	键盘输入值
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < 3; i++ {
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
	//	四周数据清0
	for i := 0; i < 3; i++ {
		// 第一行和最后一行的所有元素
		if i == 0 || i == 2 {
			for j := 0; j < 4; j++ {
				matrix[i][j] = 0
			}
		} else {
			// 中间行的第一列和最后一列
			matrix[i][0] = 0
			matrix[i][3] = 0
		}
	}
	fmt.Println("处理后的二维切片")
	for _, row := range matrix {
		fmt.Println(row)
	}

	/*
			创建了一个三行四列的二维切片matrix。然后，我们使用bufio.NewReader从标准输入读取用户输入，
		并使用ReadString函数读取一行字符串，并通过strings.TrimSpace去除字符串前后的空白字符。接着，
		我们使用strconv.Atoi将字符串转换为整数并存储在二维切片中。

			完成输入后，我们遍历二维切片，将第一行和最后一行的所有元素以及中间行的第一列和最后一列的元素设置为0。
		最后，我们打印出处理后的二维切片。
	*/
}
