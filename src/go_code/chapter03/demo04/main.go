package main
import "fmt"

// 变量注意事项
func main()  {
	// 值可以在同一类型范围内不断变化
	var i int =10
	i=30
	i=50
    fmt.Println("i=",i)
	// i=1.2 //int,不能改变数据类型
}