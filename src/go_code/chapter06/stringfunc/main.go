package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	//  1.统计字符串的长度，按字节 len(str)
	str := "hello哈" //（ascii的字符（字母和数字）占一个字节，汉字占用3个字节）
	fmt.Println("str len=", len(str))

	//  2.字符串遍历，同时处理有中文的问题转切片， r:=[]rune(str2)
	str2 := "hello哈哈哈"
	r := []rune(str2)
	for i := 0; i < len(r); i++ {
		fmt.Printf("字符=%c\n", r[i])
	}

	//	3.字符串转整数：n,err=strconv.Atoi("12")
	n, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("转换错误", err)
	} else {
		fmt.Println("转成结果是", n)
	}

	//	4.整数转字符串  str=strconv.Itoa(1234)
	strNum := strconv.Itoa(12345)
	fmt.Printf("str=%v,str=%T\n", strNum, strNum)

	//  5.字符串 转 []byte: var bytes=[]byte("hello go)
	var byteString = []byte("hello go")
	fmt.Println("byteString=", byteString)

	//	6.[]byte 转字符串：str=string([]byte{97,98,99})
	strBytes := string([]byte{97, 98, 99})
	fmt.Println("strBytes=", strBytes)

	//	7.10进制转2，8，16进制：str=strconv.FormatInt(123,2),返回对应的字符串
	str = strconv.FormatInt(123, 2)
	fmt.Printf("123对应的二进制是%v\n", str)
	str = strconv.FormatInt(123, 16)
	fmt.Printf("123对应的十六进制是%v\n", str)

	//	8.查找子字符串是否在指定的字符串中：strings.Contains("seafood","foo") //true
	b := strings.Contains("seafood", "foo") //true
	fmt.Printf("b=%v\n", b)

	//	9.统计一个字符串中有几个指定的字串：strings.Count("ceheese","e")  //4
	num := strings.Count("ceheese", "e")
	fmt.Printf("num=%v\n", num)

	//	10.不区分大小写的字符串比较（==是区分字母大小写的）：fmt.Println(strings.EqualFold("abc","ABC")) //true
	b = strings.EqualFold("abc", "ABC")
	fmt.Printf("b=%v\n", b)
	fmt.Println("结果：", "abc" == "ABC") //false

	//	11.返回子串在字符串第一次出现的index值，如果没有返回-1：strings.Index("NLT_abc","abc")//4
	index := strings.Index("NLT_abc", "abc") //4
	fmt.Println("index=", index)

	//	12.返回子串在字符串最后一次出现的index,如果没有返回-1：strings.LastIndex("go golang","go")
	index = strings.LastIndex("go golang", "go")
	fmt.Printf("index=%v\n", index)

	//	13.将指定的子串替换成另外一个子串：strings.Replace("go go hello","go","go语言",n) n可以指定你希望替换几个，如果n=-1表示全部替换
	str = strings.Replace("go go hello", "go", "哈哈哈", 1)
	fmt.Printf("str=%v\n", str)

	//	14.按照指定的某个字符，为分割标识，将一个字符串拆分成字符串数组：strings.Split("hello,world,ok",",")
	strArray := strings.Split("hello,world,ok", ",")
	for i := 0; i < len(strArray); i++ {
		fmt.Printf("str[%v]=%v\n", i, strArray[i])
	}
	fmt.Printf("strArray=%v\n", strArray)

	//	15.将字符串的字母进行大小写的转换：strings.ToLower("Go") //go  strings.ToUpper("Go") //GO
	str = "goLang Hello"
	str = strings.ToLower(str)
	fmt.Printf("str=%v\n", str)
	str = strings.ToUpper(str)
	fmt.Printf("str=%v\n", str)

	//	16.将字符串左右两边的空格去掉：strings.TrimSpace(" tn a lone gopher ntm  ")
	str = strings.TrimSpace(" tn a lone gopher ntm  ")
	fmt.Printf("str=%v\n", str)

	//	17.将字符串左右两边指定的字符去掉：strings.Trim("! hello! "," !") //["hello"]  //将左右两边!和""去掉
	str = strings.Trim("! hello! ", " !")
	fmt.Printf("str=%v\n", str)

	//	18.将字符串左边指定的字符去掉：strings.TrimLeft("! hello! "," !") //["hello"] //将左边!和""去掉
	//	19.将字符串右边指定的字符去掉：strings.TrimRight("! hello! "," !") //["hello"] //将右边!和""去掉
	//	20.判断字符串是否以指定的字符串开头：strings.HasPrefix("ftp://192.168.10.1","ftp")  //true
	b = strings.HasPrefix("ftp://192.168.10.1", "ftp")
	fmt.Printf("b=%v\n", b)
	//	21.判断字符串是否以指定的字符串结束：strings.HasSuffix("ftp://192.168.10.1","ftp")  //false
	b = strings.HasSuffix("ftp://192.168.10.1", "ftp")
	fmt.Printf("b=%v\n", b)
}
