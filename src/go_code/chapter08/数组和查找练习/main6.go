package main

import "fmt"

/*
实现查找的核心代码，例如已知数组arr[10]string里面保存了十个元素，
现在要查找"AA"在其中是否存在，打印提示，如果又多个”AA“，也要找到对应的下标
*/
func findStringIndices(arr []string, target string) []int {
	var indices []int
	for i, str := range arr {
		if str == target {
			indices = append(indices, i)
		}
	}
	return indices
}
func main() {
	arr := []string{"AA", "BB", "CC", "AA", "DD", "AA", "EE", "FF", "GG", "HH"}
	target := "AA"

	indices := findStringIndices(arr, target)
	if len(indices) > 0 {
		fmt.Printf("找到 \"%s\" 在indices里: ", target)
		for _, idx := range indices {
			fmt.Printf("%d", idx)
		}
		fmt.Println()
	} else {
		fmt.Printf("\"%s\" 数组里没找到.\n", target)
	}
}
