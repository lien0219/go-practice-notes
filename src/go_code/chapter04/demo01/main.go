package main

import "fmt"

// /  、  % 、++ 、-- 的使用
func main() {

	//   /的使用
	//	如果运算的数是整数，那么除后，去掉小数部分，保留整数部分
	fmt.Println(10 / 4)
	var n1 float32 = 10 / 4
	fmt.Println(n1)
	//	如果希望保留小数部分，则需要浮点数参与运算
	var n2 float32 = 10.0 / 4
	fmt.Println(n2)

	//	  %的使用
	// 公式：  a % b = a - a / b * b
	fmt.Println("10%3=", 10%3)     //1
	fmt.Println("-10%3=", -10%3)   //-10 - (-10) / 3 * 3=-10-(-9)=-1
	fmt.Println("10%-3=", 10%-3)   //1
	fmt.Println("-10%-3=", -10%-3) //-1

	//	 ++ 和 -- 的使用
	var i int = 10
	i++                  //等价于i=i+1
	fmt.Println("i=", i) //11
	i--                  //等价于i=i-1
	fmt.Println("i=", i) //10
}
