package main

import "fmt"

/*
1.编写一个函数就收 n int
2.能够将斐波那契的数列放到切片中
3.斐波那契数列形式：
arr[0] =1,arr[1] =1,arr[2] =2,arr[3] =3,arr[4] =5,arr[5]=8
*/

func fbn(n int) []uint64 {
	//	声明一个切片，切片大小是n
	fbnSlice := make([]uint64, n)
	fbnSlice[0] = 1
	fbnSlice[1] = 1
	for i := 2; i < n; i++ {
		fbnSlice[i] = fbnSlice[i-1] + fbnSlice[i-2]
	}
	return fbnSlice
}

func main() {
	fnbSlice := fbn(10)
	fmt.Println("fnbSlice=", fnbSlice)
}
