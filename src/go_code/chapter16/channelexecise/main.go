package main

import "fmt"

type Cat struct {
	Name string
	Age  int
}

func main() {
	allChan := make(chan interface{}, 3)

	allChan <- 10
	allChan <- "tom jack"
	cat := Cat{"狗", 4}
	allChan <- cat

	//	获取管道中的第三个元素，需先将前两个推出
	<-allChan
	<-allChan

	newCat := <-allChan

	fmt.Printf("newCat=%T，newCat=%v \n", newCat, newCat)
	//类型断言
	a := newCat.(Cat)
	fmt.Printf("newCat.Name=%v", a.Name)

}
