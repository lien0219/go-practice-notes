package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"sync"
	"time"
)

/*
开一个协程writeDataToFile,随机生成1000个数据，存放到文件中
当writeDataToFile完成写1000个数据到文件后，让sort协程从文件中读取1000个文件
并完成排序，重新写入到另外一个文件
开10个协程writeDataToFile,每个协程随机生成1000个数据，存放到10文件中
当十个文件都生成了，让10个sort协程从10个文件中读取1000个文件并
完成排序，重新写入到十个结果文件
*/

// 写入随机数据到文件
func writeDataToFile(filename string, wg *sync.WaitGroup) {
	defer wg.Done()
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for i := 0; i < 1000; i++ {
		_, err := writer.WriteString(fmt.Sprintf("%d\n", rand.Intn(10000))) // 生成0到9999的随机数
		if err != nil {
			panic(err)
		}
	}
	writer.Flush()
}

// 从文件读取数据，排序后重新写入文件
func sortDataFromFile(srcFilename, dstFilename string, wg *sync.WaitGroup) {
	defer wg.Done()

	var numbers []int
	file, err := os.Open(srcFilename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var num int
		_, err := fmt.Sscanf(scanner.Text(), "%d", &num)
		if err != nil {
			continue // 忽略无法解析为整数的行
		}
		numbers = append(numbers, num) // 将解析的整数添加到切片中
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	sort.Ints(numbers)

	outputFile, err := os.Create(dstFilename)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	for _, num := range numbers {
		_, err := outputFile.WriteString(fmt.Sprintf("%d\n", num))
		if err != nil {
			panic(err)
		}
	}
}
func main() {
	rand.Seed(time.Now().UnixNano())
	var wgWrite, wgSort sync.WaitGroup

	//	写入数据到十个文件
	for i := 1; i <= 10; i++ {
		wgWrite.Add(1)
		go writeDataToFile(fmt.Sprintf("data%d.txt", i), &wgWrite)
	}
	wgWrite.Wait() //等待所有写操作完成

	//	排序写入到10个新文件
	for i := 1; i <= 10; i++ {
		wgSort.Add(1)
		go sortDataFromFile(fmt.Sprintf("data%d.txt", i), fmt.Sprintf("sorted_data%d.txt", i), &wgSort)
	}
	wgSort.Wait() // 等待所有排序和写入操作完成
	fmt.Println("所有文件都已处理和排序完成。")
}
