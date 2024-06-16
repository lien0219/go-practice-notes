package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// 1.声明hero结构体
type Hero struct {
	Name string
	Age  int
}

// 2.声明一个hero结构体切片类型
type HeroSlice []Hero

// 3.实现interface 接口
func (hs HeroSlice) Len() int {
	return len(hs)
}

// Less方法就是决定你使用什么标准进行排序
// 1.按照Hero得年龄从小到大排序
func (hs HeroSlice) Less(i, j int) bool {
	//按照年龄排序
	return hs[i].Age < hs[j].Age
	//	按照名字排序
	//return hs[i].Name < hs[j].Name
}
func (hs HeroSlice) Swap(i, j int) {
	//交换
	//temp := hs[i]
	//hs[i] = hs[j]
	//hs[j] = temp
	//等价于上
	hs[i], hs[j] = hs[j], hs[i]
}

// 1.声明一个Student结构体
type Student struct {
	Name  string
	Age   int
	Score float64
}

func main() {
	var intSlice = []int{0, -1, 10, 7, 90}
	//	对intSlice切片进行排序
	sort.Ints(intSlice)
	fmt.Println(intSlice)

	//	编写一个结构体切片排序
	//	测试
	var heroes HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("英雄~%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		//	将hero append到heroes切片
		heroes = append(heroes, hero)
	}
	//	排序前顺序
	for _, v := range heroes {
		fmt.Println(v)
	}
	//	排序后
	sort.Sort(heroes)
	fmt.Println("——————————————————排序后————————————————")
	for _, v := range heroes {
		fmt.Println(v)
	}
}
