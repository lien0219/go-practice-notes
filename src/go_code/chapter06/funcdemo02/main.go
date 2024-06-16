package main

import "fmt"

func test(n1 int) {
	n1 = n1 + 1
	fmt.Println("test() n1=", n1)
}

// 返回值
func getNum(n1 int, n2 int) int {
	sum := n1 + n2
	fmt.Println("getNum sum=", sum)
	return sum
}

// 返回多个值
func getSumAndSub(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}

func main() {
	n1 := 10
	test(n1)
	fmt.Println("main() n1=", n1)

	sum := getNum(10, 20)
	fmt.Println("main() sum=", sum)

	res1, res2 := getSumAndSub(1, 2)
	fmt.Printf("res1=%v res2=%v\n", res1, res2)

	_, res3 := getSumAndSub(2, 1)
	fmt.Printf("res3=%v", res3)
}
