package main
import "fmt"

func main()  {
	//1. 指定变量类型，声明后但不赋值，则使用默认值（int类型的默认值是0）
	var i int
	fmt.Println("i=",i)

	// 2. 根据值自行判断变量类型（类型推导）
	var num=10.11
	fmt.Println("num=",num)

	// 3. 省略var,:=左侧的变量不应该是已经声明过的，否则会报错
	// 等价于var name string     name='tom
	//  := 的:不能省略，否则错误
	name := "tom"
	fmt.Println("name=",name)
}