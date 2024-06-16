package main

import (
	"fmt"
	"time"
)

func main() {

	//	日期和时间相关的函数和方法使用

	//	1.获取当前的时间
	now := time.Now()
	fmt.Printf("now=%v now type=%T\n", now, now)

	//	2.通过now获取年月日、时分秒
	fmt.Printf("年=%v\n", now.Year())
	fmt.Printf("月=%v\n", int(now.Month()))
	fmt.Printf("日=%v\n", now.Day())
	fmt.Printf("时=%v\n", now.Hour())
	fmt.Printf("分=%v\n", now.Minute())
	fmt.Printf("秒=%v\n", now.Second())
	//	格式化日期时间
	fmt.Printf("当前年月日 %02d-%02d-%02d %02d:%02d:%02d \n", now.Year(), now.Month(),
		now.Day(), now.Hour(), now.Minute(), now.Second())
	//  格式化日期时间第二种方式
	fmt.Printf(now.Format("2006/01-02 15:04:05"))
	fmt.Println()
	fmt.Printf(now.Format("2006-01-02"))
	fmt.Println()
	fmt.Printf(now.Format("15:04:05"))
	fmt.Println()

	fmt.Printf(now.Format("01")) //月份
	fmt.Println()

	//6.结合Sleep使用时间常量
	//Nanosecond  Duration = 1
	//Microsecond          = 1000 * Nanosecond
	//Millisecond          = 1000 * Microsecond
	//Second               = 1000 * Millisecond
	//Minute               = 60 * Second
	//Hour                 = 60 * Minute
	//	每隔1秒打印一个数字，打印到100退出
	i := 0
	for {
		i++
		fmt.Println(i)
		//	休眠
		time.Sleep(time.Second) //一秒
		//time.Sleep(time.Millisecond * 100) //0.1秒
		if i == 2 {
			break
		}
	}

	//	7.Unix和UnixNano的使用（时间戳，前者是秒，后者是纳秒）
	fmt.Printf("unix时间戳=%v unixnano时间戳=%v\n", now.Unix(), now.UnixNano())
}
