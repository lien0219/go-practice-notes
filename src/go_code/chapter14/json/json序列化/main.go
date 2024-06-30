package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name     string  `json:"monster_name"`
	Age      int     `json:"monster_age"`
	Birthday string  `json:"monster_birthday"`
	Sal      float64 `json:"monster_sal"`
	Skill    string  `json:"monster_skill"`
}

// 结构体序列化
func testStruct() {
	monster := Monster{
		Name:     "牛头",
		Age:      500,
		Birthday: "2024-1-1",
		Sal:      8000.0,
		Skill:    "降龙十八掌",
	}
	//	monster序列化
	//	func Marshal(v interface{}) ([]byte, error)
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Printf("序列号错误 err=%v \n", err)
	}
	fmt.Printf("monster序列化后：%v \n", string(data))
}

// Map序列化
func testMap() {
	var a map[string]interface{}
	a = make(map[string]interface{})
	a["name"] = "张三"
	a["age"] = 30
	a["address"] = "洪崖洞"

	//	将map进行序列化
	data, err := json.Marshal(a)
	if err != nil {
		fmt.Printf("序列号错误 err=%v \n", err)
	}
	fmt.Printf("a map序列化后：%v \n", string(data))
}

// 切片序列化
func testSlice() {
	var slice []map[string]interface{}

	var m1 map[string]interface{}
	m1 = make(map[string]interface{})
	m1["name"] = "jack"
	m1["age"] = "7"
	m1["address"] = "北京"
	slice = append(slice, m1)

	var m2 map[string]interface{}
	m2 = make(map[string]interface{})
	m2["name"] = "tom"
	m2["age"] = "20"
	m2["address"] = [2]string{"美国", "夏威夷"}
	slice = append(slice, m2)

	//	切片序列化
	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Printf("序列号错误 err=%v \n", err)
	}
	fmt.Printf("slice序列化后：%v \n", string(data))
}

// 基本类型序列化 (基本数据类型序列化没意义)
func testFloat64() {
	var num1 float64 = 2345.67

	data, err := json.Marshal(num1)
	if err != nil {
		fmt.Printf("序列号错误 err=%v \n", err)
	}
	fmt.Printf("num1序列化后：%v \n", string(data))
}

func main() {
	//	演示结构体、map、切片进行序列化
	testStruct()
	testMap()
	testSlice()
	testFloat64() //基本数据类型序列化
}
