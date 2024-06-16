package main

import "fmt"

func main() {
	var a map[string]string
	//	使用map前需要先make,make的作用是给map分配数据空间
	a = make(map[string]string, 10)
	a["no1"] = "张三"
	a["no2"] = "李四"
	a["no3"] = "王五"
	a["no1"] = "哈哈哈"
	fmt.Println(a)

	//	第二种方式
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "上海"
	cities["no3"] = "天津"
	fmt.Println(cities)

	//	第三种方式
	heroes := map[string]string{
		"hero1": "松江",
		"hero2": "我",
	}
	fmt.Println(heroes)
}
