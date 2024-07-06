package monster

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Monster struct {
	Name  string
	Age   int
	Skill string
}

func (this *Monster) Store() bool {
	//	序列化
	data, err := json.Marshal(this)
	if err != nil {
		fmt.Println("marshal err:", err)
		return false
	}
	//	保存到文件
	filePath := "D:\\桌面\\技术栈\\Golang\\go_project\\src\\go_code\\chapter15\\testcase02\\monsterFile.ser"
	err = ioutil.WriteFile(filePath, data, 0666)
	if err != nil {
		fmt.Println("write file err:", err)
		return false
	}
	return true
}

func (this *Monster) ReStore() bool {
	//	读取序列化的字符串
	filePath := "D:\\桌面\\技术栈\\Golang\\go_project\\src\\go_code\\chapter15\\testcase02\\monsterFile.ser"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("read file err:", err)
		return false
	}
	//	反序列化
	err = json.Unmarshal(data, this)
	if err != nil {
		fmt.Println("Unmarshal err:", err)
		return false
	}
	return true
}
