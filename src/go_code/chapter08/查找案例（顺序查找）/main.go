package main

import "fmt"

/*
有一个数列：狮子，老虎，公鸡，猴子
从键盘中输入任意一个名称，判断数列中是否包含此名称（顺序查找）
*/
func main() {
	names := [4]string{"狮子", "老虎", "公鸡", "猴子"}
	var animals = ""
	fmt.Println("请输入要查找的动物名！")
	fmt.Scanln(&animals)
	//	顺序查找：第一种方式
	//for i := 0; i < len(names); i++ {
	//	if animals == names[i] {
	//		fmt.Printf("找到%v,下标%v", animals, i)
	//		break
	//	} else if i == (len(names) - 1) {
	//		fmt.Printf("没有找到%v \n", animals)
	//	}
	//}

	//	顺序查找：第二种方式
	index := -1
	for i := 0; i < len(names); i++ {
		if animals == names[i] {
			index = 1
			break
		}
	}
	if index != -1 {
		fmt.Printf("找到%v,下标%v", animals, index)
	} else {
		fmt.Printf("没有找到%v \n", animals)
	}
}
