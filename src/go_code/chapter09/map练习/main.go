package main

import "fmt"

/*
演示key-value 的value是map的案例

存放三个学生信息，每个学生有name和sex信息
*/
func main() {
	students := make(map[string]map[string]string)
	students["stu01"] = make(map[string]string, 3)
	students["stu01"]["name"] = "tom"
	students["stu01"]["sex"] = "男"
	students["stu01"]["address"] = "北京"

	students["stu02"] = make(map[string]string, 3)
	students["stu02"]["name"] = "mary"
	students["stu02"]["sex"] = "女"
	students["stu02"]["address"] = "上海"

	fmt.Println(students)
}
