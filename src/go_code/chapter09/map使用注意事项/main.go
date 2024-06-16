package main

import "fmt"

func modify(map1 map[int]int) {
	map1[10] = 900
}

// 学生结构体
type Stu struct {
	Name    string
	Age     int
	Address string
}

func main() {

	//	map是引用类型，遵守引用类型传递机制，在一个函数接收map
	//	修改后会直接修改原来的map
	map1 := make(map[int]int, 2) //超出容量后map会自己增加不会报错
	map1[1] = 90
	map1[2] = 88
	map1[10] = 1
	map1[20] = 2
	modify(map1)
	fmt.Println(map1)

	//	map的value经常使用struct类型
	//	更适合管理复杂的数据
	//	例如value为Student的结构体
	//	1.map的key为学生的学号，是唯一的
	//	2.map的value为结构体，包含学生、姓名、年龄、地址
	students := make(map[string]Stu, 10)
	//创建两个学生
	stu1 := Stu{"tom", 18, "河南"}
	stu2 := Stu{"mary", 28, "北京"}
	students["no1"] = stu1
	students["no2"] = stu2
	fmt.Println(students)

	//	遍历学生
	for k, v := range students {
		fmt.Printf("学生的编号是%v \n", k)
		fmt.Printf("学生的名字是%v \n", v.Name)
		fmt.Printf("学生的年龄是%v \n", v.Age)
		fmt.Printf("学生的地址是%v \n", v.Address)
		fmt.Println()
	}
}
