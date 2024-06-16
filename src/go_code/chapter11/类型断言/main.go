package main

import "fmt"

type Point struct {
	x int
	y int
}

func main() {
	var a interface{}
	var point Point = Point{1, 2}
	a = point
	var b Point
	//b=a不可以
	//b=a.(Point) 可以
	b = a.(Point) //类型断言
	fmt.Println(b)

	//	类型断言的其他案例
	//var x interface{}
	//var b2 float32 = 1.1
	//x = b2
	//y := x.(float32)
	/*
		注意：
			在进行类型断言时，如果类型不匹配，就会报panic,因此进行
			类型断言时，需要确保原来的空接口指向的就是断言的类型
	*/
	//fmt.Printf("y的类型：%T 值：%v", y, y)

	//	类型断言（带检测的）
	var x interface{}
	var b2 float32 = 1.1
	x = b2
	//	类型断言带检测
	y, ok := x.(float32)
	//也可以这样写
	//if y, ok := x.(float32);ok {
	//	fmt.Println("转换成功")
	//	fmt.Printf("y的类型：%T 值：%v", y, y)
	//}else {
	//	fmt.Println("转换失败")
	//}
	if ok == true {
		fmt.Println("转换成功")
		fmt.Printf("y的类型：%T 值：%v", y, y)
	} else {
		fmt.Println("转换失败")
	}
	fmt.Println("继续执行。。。")
}
