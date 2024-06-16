package main

import "fmt"

/*
	1.使用map[string]map[string]string的map类型
	2.key表示用户名，是唯一的，不可重复
	3.如果某个用户名存在，就将其密码修改“888888”，如果不存在就增加这个用户信息（包括昵称nickname和密码pwd）
	4.编写函数modifyUser(users map[string]map[string]string,name string)完成描述功能

*/

func modifyUser(users map[string]map[string]string, name string) {
	//v,ok:=users[name]
	if users[name] != nil {
		users[name]["pwd"] = "888888"
	} else {
		users[name] = make(map[string]string, 2)
		users[name]["pwd"] = "888888"
		users[name]["nickname"] = "昵称~" + name
	}
}

func main() {

	users := make(map[string]map[string]string, 10)
	users["smith"] = make(map[string]string, 2)
	users["smith"]["pwd"] = "999999"
	users["smith"]["nickname"] = "小猫"

	modifyUser(users, "tom")
	modifyUser(users, "mary")
	modifyUser(users, "smith")
	fmt.Println(users)
}
