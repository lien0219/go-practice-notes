package main
import "fmt"

// 全局变量
var n1=100
var n2=200
var name="jack"

// 一次性声明
var(
	n3=300
	n4=900
	name2="lien"
)

func main()  {
	// 声明多个变量

	// 方式1
	// var n1,n2,n3 int
	// fmt.Println("n1=",n1,"n2=",n2,"n3=",n3)

	// 方式2
	// var n1,name,n3=100,"tom",888
	// fmt.Println("n1=",n1,"name=",name,"n3=",n3)

	// 方式3（类型推导）
	//  n1,name,n3 :=100,"tom~",888
	// fmt.Println("n1=",n1,"name=",name,"n3=",n3)


	// 全局变量
	fmt.Println("n3=",n3,"name2=",name2,"n4=",n4)
}