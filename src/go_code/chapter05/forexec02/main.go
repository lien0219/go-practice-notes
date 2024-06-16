package main

import "fmt"

// for练习
func main() {
	/*
		统计三个班的成绩情况，每个班有五名同学
		求出各个班的平均分和所有班级的平均分（学生成绩从键盘输入）
		统计三个班及格人数
	*/

	var classNum int = 2       //班级个数
	var stuNum int = 5         //一个班级人数
	var totalSum float64 = 0.0 //所有班级的平均分
	var passCount int = 0      //及格人数

	for j := 1; j <= classNum; j++ {
		sum := 0.0 //各个班的平均分
		for i := 1; i <= stuNum; i++ {
			var score float64
			fmt.Printf("请输入第%d班 第%d个学生的成绩\n", j, i)
			fmt.Scanln(&score)
			//	累计总分
			sum += score
			//是否及格
			if score > 60 {
				passCount++
			}
		}
		fmt.Printf("第%d班级的平均分是%v\n", j, sum/5)
		totalSum += sum
	}
	fmt.Printf("各个班级的总成绩%v 所有班级平均分是%v\n", totalSum, totalSum/float64(stuNum*classNum))
	fmt.Printf("及格人数为%v\n", passCount)
}
