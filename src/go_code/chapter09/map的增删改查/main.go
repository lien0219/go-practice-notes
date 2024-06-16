package main

import "fmt"

func main() {

	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "上海"
	cities["no3"] = "天津"
	fmt.Println(cities)

	//因为no3这个key已经存在，因此下面的语句就是修改值
	cities["no3"] = "天津~~~"
	fmt.Println(cities)

	//删除操作
	delete(cities, "no1")
	fmt.Println(cities)

	//	map查找
	val, ok := cities["no2"] // 如果键不存在，ok 的值为 false，val 的值为该类型的零值
	if ok {
		fmt.Printf("no2 key 值为%v\n", val)
	} else {
		fmt.Printf("没有no2 key\n")
	}

	//	如果想一次性删除所有key
	//	1.遍历所有的key逐一删除
	//	2.直接make一个新空间
	cities = make(map[string]string)
	fmt.Println(cities)
}
