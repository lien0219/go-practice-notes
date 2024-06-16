package main

import "fmt"

/*
定义二维数组，用于保存三个班，每个班五名同学成绩
求出每个班平均分以及所有班平均分
*/
func main() {
	var scores [3][5]float64
	for i := 0; i < len(scores); i++ {
		for j := 0; j < len(scores[i]); j++ {
			fmt.Printf("请输入第%d班的第%d个学生的成绩\n", i+1, j+1)
			fmt.Scanln(&scores[i][j])
		}
	}
	//fmt.Println(scores)
	totalSum := 0.0 //所有班级的总分
	for i := 0; i < len(scores); i++ {
		sum := 0.0 //各个班级的总分
		for j := 0; j < len(scores[i]); j++ {
			sum += scores[i][j]
		}
		totalSum += sum
		fmt.Printf("第%d班级的总分为%v,平均分%v\n", i+1, sum, sum/float64(len(scores[i])))
	}
	fmt.Printf("所有班级的总分为%v,所有班级的平均分%v\n", totalSum, totalSum/15)
}
