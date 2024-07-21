package main

import (
	"fmt"
	"reflect"
)

func reflectTest01(b interface{}) {
	//	通过反射获取传入变量的type、kind、值
	rType := reflect.TypeOf(b)
	fmt.Println("rtype =", rType)

	//获取reflect.Value
	rVal := reflect.ValueOf(b)

	n2 := 2 + rVal.Int()
	fmt.Println("n2 =", n2)

	fmt.Printf("rval =%v rVal type=%T \n", rVal, rVal)

	//	将rVal转成interface{}
	iV := rVal.Interface()
	//	将interface{}通过断言转成需要的类型
	num2 := iV.(int)
	fmt.Println("num2=", num2)
}

func reflectTest02(b interface{}) {
	//	通过反射获取传入变量的type、kind(变量类别)、值
	rType := reflect.TypeOf(b)
	fmt.Println("rtype =", rType)

	//获取reflect.Value
	rVal := reflect.ValueOf(b)

	//	获取变量对应的Kind
	//1.rVal.Kind()
	kind1 := rVal.Kind()
	//2.rType.Kind()
	kind2 := rType.Kind()
	fmt.Printf("kind=%v kind=%v\n", kind1, kind2)

	//	将rVal转成interface{}
	iV := rVal.Interface()
	fmt.Printf("iv=%v iv type=%T\n", iV, iV)
	//	将interface{}通过断言转成需要的类型
	stu, ok := iV.(Student)
	if ok {
		fmt.Printf("stu.Name=%v\n", stu.Name)
	}
}

type Student struct {
	Name string
	Age  int
}
type Monster struct {
	Name string
	Age  int
}

func main() {

	//	演示对（基本数据类型、interface{}、reflect.Value）进行反射的基本操作
	var num int = 100
	reflectTest01(num)

	stu := Student{
		Name: "tom",
		Age:  123,
	}
	reflectTest02(stu)
}
