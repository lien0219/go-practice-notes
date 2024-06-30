package main

import (
	"encoding/json"
	"fmt"
)

// 将json字符串反序列化为struct
type Monster struct {
	Name     string  `json:"monster_name"`
	Age      int     `json:"monster_age"`
	Birthday string  `json:"monster_birthday"`
	Sal      float64 `json:"monster_sal"`
	Skill    string  `json:"monster_skill"`
}

func unmarshalStruct() {
	str := "{\"monster_name\":\"牛头\",\"monster_age\":500,\"monster_birthday\":\"2024-1-1\",\"monster_sal\":8000,\"monster_skill\":\"降龙八掌\"}\n"

	//	Monster实例
	var monster Monster
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Printf("unmarshal err=%v \n", err)
	}
	fmt.Printf("结构体反序列化后：%v\n", monster)
}

// 将json字符串反序列化为map
func unmarshalMap() {
	str := "{\"address\":\"洪崖洞\",\"age\":30,\"name\":\"张三\"}"

	var a map[string]interface{}
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Printf("unmarshal err=%v \n", err)
	}
	fmt.Printf("map反序列化后：%v\n", a)
}

// 将json反序列化为切片
func unmarshalSlice() {
	str := "[{\"address\":\"北\uE000京\uE000\",\"age\":\"7\",\"name\":\"jack\"},{\"address\":[\"美\uE000国\uE000\",\"夏\uE000威\uE000夷\uE000\"],\"age\":\"20\",\"name\":\"tom\"}]"

	var slice []map[string]interface{}
	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		fmt.Printf("unmarshal err=%v \n", err)
	}
	fmt.Printf("切片反序列化后：%v\n", slice)
}
func main() {
	unmarshalStruct()
	unmarshalMap()
	unmarshalSlice()
}
