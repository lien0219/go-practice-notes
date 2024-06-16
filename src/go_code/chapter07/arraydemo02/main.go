package main

import "fmt"

func main() {
	/*
		    var intArr [3]int
			//	定义完数组后，数组里的每个元素都有默认值0
			fmt.Println(intArr)

			//	从终端循环输入5个成绩，保存到float64数组并输出
			var score [5]float64
			for i := 0; i < len(score); i++ {
				fmt.Println("请输入第%d个元素得值\n", i+1)
				fmt.Scanln(&score[i])
			}
			//变量数组打印
			for i := 0; i < len(score); i++ {
				fmt.Printf("score[%d]=%v", i, score[i])
			}
	*/

	//	四种初始化数组的方式
	//	方式1
	var numArr01 [3]int = [3]int{1, 2, 3}
	fmt.Println("numArr01=", numArr01)
	//	方式2
	var numArr02 = [3]int{6, 7, 8}
	fmt.Println("numArr02=", numArr02)
	//	方式3  不指定长度
	var numArr03 = [...]int{11, 22, 33}
	fmt.Println("numArr03=", numArr03)
	//	方式4  指定对应值
	var numArr04 = [...]int{1: 100, 0: 200, 3: 300}
	fmt.Println("numArr04=", numArr04)
	//	类型推导
	strArr := [...]string{1: "tom", 0: "lisi", 2: "tut"}
	fmt.Println("strArr=", strArr)
}
