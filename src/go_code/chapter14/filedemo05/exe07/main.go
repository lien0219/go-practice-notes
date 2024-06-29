package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 6.统计英文、数字、空格和其他字符数量
type CharCount struct {
	ChCount    int //英文个数
	NumCount   int //数字个数
	SpaceCount int //空格个数
	OtherCount int //其他字符个数
}

func main() {

	//	打开文件，读取行统计，结果保存结构体
	fileName := "D:\\桌面\\技术栈\\Golang\\go_project\\src\\go_code\\chapter14\\filedemo05\\file\\tongji.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}
	defer file.Close()
	var count CharCount
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF { //文件末尾退出
			break
		}
		//	遍历str,进行统计
		for _, v := range str {
			//fmt.Println(v)
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough //穿透处理
			case v >= 'A' && v <= 'Z':
				count.ChCount++
			case v == ' ' || v == '\t':
				count.SpaceCount++
			case v >= '0' && v <= '9':
				count.NumCount++
			default:
				count.OtherCount++
			}
			//if (v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z') {
			//	count.ChCount++
			//} else if v == ' ' || v == '\t' {
			//	count.SpaceCount++
			//} else if v >= '0' && v <= '9' {
			//	count.NumCount++
			//} else {
			//	count.OtherCount++
			//}
		}
	}
	fmt.Printf("字符个数：%v 数字个数：%v 空格个数：%v 其他字符个数：%v",
		count.ChCount, count.NumCount, count.SpaceCount, count.OtherCount)
}
