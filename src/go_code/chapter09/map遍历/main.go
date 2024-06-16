package main

import "fmt"

func main() {
	//	for-range遍历
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "上海"
	cities["no3"] = "深圳"
	for k, v := range cities {
		fmt.Printf("k=%v v=%v\n", k, v)
	}
	fmt.Println("cities的长度：", len(cities))

	//for-range遍历复杂map
	students := make(map[string]map[string]string)
	students["stu01"] = make(map[string]string, 3)
	students["stu01"]["name"] = "tom"
	students["stu01"]["sex"] = "男"
	students["stu01"]["address"] = "北京"

	students["stu02"] = make(map[string]string, 3)
	students["stu02"]["name"] = "mary"
	students["stu02"]["sex"] = "女"
	students["stu02"]["address"] = "上海"

	for k1, v1 := range students {
		fmt.Println("k1=", k1)
		for k2, v2 := range v1 {
			fmt.Printf("\t k2=%v v2=%v\n", k2, v2)
		}
		fmt.Println()
	}
}
