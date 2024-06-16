package main

import "fmt"

/*
问题：
编写一个函数，输出100以内所有的素数，每行显示五个并求和
*/
/*
思路：
定义了一个isPrime函数来判断一个数是否为素数。
然后，printPrimes函数遍历2到99的所有整数，
检查它们是否为素数，并计算它们的和。
每当找到五个素数时，就打印一个新行。
最后，在main函数中调用printPrimes函数并打印素数的和。
*/

// 判断一个数是否为素数
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// 打印100以内所有的素数，每行五个，并返回所有素数的和
func printPrimes() int {
	sum := 0
	count := 0
	for i := 2; i < 100; i++ {
		if isPrime(i) {
			sum += i
			count++
			fmt.Printf("数字=%d", i)
			if count%5 == 0 {
				fmt.Println()
			}
		}
	}
	// 如果最后一个素数不是每行的第五个，则换行
	if count%5 != 0 {
		fmt.Println()
	}
	return sum
}

func main() {
	sum := printPrimes()
	fmt.Printf("100以内所有素数的和是: %d\n", sum)
}
