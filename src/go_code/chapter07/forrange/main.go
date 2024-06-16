package main

import "fmt"

func main() {
	/*
		现在有六只鸡，它们体重分别是3kg\5kg\1kg\3.4kg\3kg\50kg
		请问六只鸡的总体重是多少？平均体重是多少？
	*/

	var hens [6]float64
	hens[0] = 3.0
	hens[1] = 5.0
	hens[2] = 1.0
	hens[3] = 3.4
	hens[4] = 2.0
	hens[5] = 50.0

	totalWeight := 0.0
	for i := 0; i < len(hens); i++ {
		totalWeight += hens[i]
	}
	avWeight := fmt.Sprintf("%.2f", totalWeight/6)
	fmt.Printf("totalWeight=%v avWeight=%v\n", totalWeight, avWeight)
}
