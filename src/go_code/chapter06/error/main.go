package main

import (
	"errors"
	"fmt"
)

// 错误处理  1.使用defer + recover 来捕获和处理异常 || 2.panic()  内置函数，输出错误信息并退出程序
func test() {
	defer func() {
		//判断方式一
		//err := recover() //recover()是内置函数，可以捕获到异常
		//if err != nil {
		//	fmt.Println("err=", err)
		//	fmt.Println("发送信息给管理员")
		//}

		//	方式二   if后面可以跟一个语句再判断
		//recover()是内置函数，可以捕获到异常
		if err := recover(); err != nil {
			fmt.Println("err=", err)
			fmt.Println("发送信息给管理员")
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res=", res)
}

// 函数读取配置文件init.conf的信息
// 如果文件名传入不正确，就返回一个自定义的错误
func readConf(name string) (err error) {
	if name == "config.ini" {
		return nil
	} else {
		//	自定义错误
		return errors.New("读取文件错误...")
	}
}

func test02() {

	err := readConf("config.ini")
	if err != nil {
		//	输出错误
		panic(err)
	}
	fmt.Println("test02()继续执行...")
}

func main() {
	//	测试 test()
	//test()
	//i := 0
	//for {
	//	i++
	//	fmt.Println("main函数下的代码输出")
	//	time.Sleep(time.Second) //休眠一秒
	//	if i > 2 {
	//		break
	//	}
	//}

	//	测试test02()
	test02()
	fmt.Println("main函数下的代码输出...")
}
