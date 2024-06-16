package main

import (
	"fmt"
	"time"
)

/*
问题：
编写一个函数来判断在1990年1月1日之后的某一天是打鱼还是晒网。
由于打鱼-晒网的周期是五天（三天打鱼，两天晒网），
使用日期的天数差来确定在给定的日期上是处于这个周期的哪个阶段。
*/
/*
思路：
使用time.Date函数来创建一个time.Time类型的日期对象，
并使用Sub方法来计算两个日期之间的时间差。
由于Sub方法返回的是一个time.Duration类型，
需要将其转换为天数。由于time.Duration是以纳秒为单位的，
我们通过除以每天的小时数（24）和每小时的秒数（3600）
以及每秒的纳秒数（1e9）来得到天数差。但是，为了简化计算，
可以直接除以每天的小时数（24）并假设所有的时间都是午夜开始计算的
（即没有小时、分钟、秒和纳秒的偏移）。

此外，请注意日期格式字符串"2006-01-02"是Go语言中
用于解析和格式化时间的预定义
引用时间Mon Jan 2 15:04:05 MST 2006的简写形式。
*/

// isFishingDay 判断给定日期是打鱼日还是晒网日
func isFishingDay(date time.Time) bool {
	// 设置起始日期为1990年1月1日
	startDate := time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)
	// 计算给定日期与起始日期之间的天数差
	daysDiff := date.Sub(startDate).Hours() / 24
	// 根据天数差和打鱼-晒网周期（5天）来确定是打鱼日还是晒网日
	// 如果余数为0、1或2，则是打鱼日（包括周期开始的第一天）
	// 否则是晒网日
	remainder := int(daysDiff) % 5
	return remainder <= 2
}

func main() {
	//判断2024年某个日期是打鱼日还是晒网日
	checkDate := time.Date(2024, 3, 15, 0, 0, 0, 0, time.UTC)
	if isFishingDay(checkDate) {
		fmt.Printf("在%s这一天是打鱼日\n", checkDate.Format("2006-01-02"))
	} else {
		fmt.Printf("在%s这一天是晒网日\n", checkDate.Format("2006-01-02"))
	}
}
