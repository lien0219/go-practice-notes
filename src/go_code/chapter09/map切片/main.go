package main

import "fmt"

func main() {
	//	map切片演示
	/*
		使用一个map来记录monster的信息name和age，也就是说一个
		monster对应一个map，并且妖怪的个数可以动态增加=>map切片
	*/
	var monsters []map[string]string
	monsters = make([]map[string]string, 2)
	if monsters[0] == nil {
		monsters[0] = make(map[string]string, 2)
		monsters[0]["name"] = "牛魔"
		monsters[0]["age"] = "500"
	}
	if monsters[1] == nil {
		monsters[1] = make(map[string]string, 2)
		monsters[1]["name"] = "玉兔"
		monsters[1]["age"] = "900"
	}

	newMonsters := map[string]string{
		"name": "新-猪八戒",
		"age":  "200",
	}
	monsters = append(monsters, newMonsters)
	fmt.Println(monsters)
}
