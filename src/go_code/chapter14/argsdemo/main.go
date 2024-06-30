package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("命令行参数：", len(os.Args))
	//	遍历os.Args切片，得到所有的命令行输入参数值
	for i, v := range os.Args {
		fmt.Printf("argsdemo[%v]=%v \n", i, v)
	}
}
