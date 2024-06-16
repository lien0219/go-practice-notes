package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

/*
练习1
实现循环打印输入的月份天数，使用continue实现，要有判断输入的月份是否错误的语句
*/
func main() {

	/*
			流程
			首先使用bufio.NewReader(os.Stdin)来创建一个读取器，以便从标准输入读取数据。
			然后，使用for循环来不断提示用户输入月份。使用ReadString('\n')读取一行输入，
			并使用strconv.Atoi将字符串转换为整数。
			接着，检查输入的月份是否在1到12之间。如果不在这个范围内，或者转换过程中出现错误，
			打印错误消息并使用continue跳过本次循环的剩余部分。
			如果月份输入正确，使用time包来获取下一个月的第一天，并减去当前月份的第一天来
			得到这个月的天数。最后，打印出月份和天数。

			注意：
			如果希望程序在打印一次后退出，可以在打印后加上break语句。如果想要程序持续循环直到用户手动停止（例如通过Ctrl+C），则不需要break语句。
		    导入strings包来使用strings.TrimSpace函数去除输入中的换行符和空格。
	*/
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("请输入月份（1-12）:")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		month, err := strconv.Atoi(text)

		if err != nil || month < 1 || month > 12 {
			fmt.Println("输入的月份不正确，请输入1到12之间的数字")
			continue // 跳过本次循环的剩余部分
		}

		// 使用time包判断月份的天数（这里使用time.Date来获取下个月的第一天，然后减去1天）
		nextMonth := time.Date(time.Now().Year(), time.Month(month)+1, 0, 0, 0, 0, 0, time.Local)
		daysInMonth := nextMonth.Sub(time.Date(time.Now().Year(), time.Month(month), 1, 0, 0, 0, 0, time.Local)).Hours() / 24
		fmt.Printf("月份 %d 有 %d 天\n", month, int(daysInMonth))
		break //打印一次退出
	}
}
