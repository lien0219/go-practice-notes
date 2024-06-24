package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// 4.将文件a文件内容导入到文件b

// 判断文件是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil { //文件或者目录存在
		return false, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
func main() {
	/*
		1.将文件a读取到内存
		2.将读取到的内容写入文件b
	*/

	file1Path := "D:\\桌面\\技术栈\\Golang\\go_project\\src\\go_code\\chapter14\\filedemo05\\file\\A.txt"
	file2Path := "D:\\桌面\\技术栈\\Golang\\go_project\\src\\go_code\\chapter14\\filedemo05\\file\\B.txt"

	cunzai, _ := PathExists(file1Path)
	fmt.Println("文件是否存在：", cunzai)

	data, err := ioutil.ReadFile(file1Path)
	if err != nil {
		fmt.Printf("read file err=%v", err)
		return
	}

	err = ioutil.WriteFile(file2Path, data, 0666)
	if err != nil {
		fmt.Printf("write file err=%v \n", err)
	}
}
