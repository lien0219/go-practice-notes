package main

import (
	"fmt"
	"reflect"
)

func main() {
	var str string = "hello world"
	fs := reflect.ValueOf(&str)
	fs.Elem().SetString("tom")
	fmt.Printf("%v\n", str)
}
