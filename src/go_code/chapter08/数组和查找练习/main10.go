package main

import (
	"fmt"
	"math"
	"sort"
)

/*
实现跳水比赛，8个评委打分。运动员的成绩是8个成绩取掉一个最高分，
去掉一个最低分，剩下的6个分数的平均分就是最后得分，使用一维数组实现如下功能：
（1）请把打最高分的评委和最低分的评委找出来。
（2）找出最佳评委和最差评委。最佳评委就是打分和最后得分最接近的评委。最差评委就是打分和最后得分相差最大的。
*/

// 最终得分
func calculateFinalScore(scores []int) (float64, error) {
	if len(scores) < 8 {
		return 0, fmt.Errorf("需要至少8个评委的分数")
	}
	//	排序分数
	sort.Ints(scores)
	// 去掉最高分和最低分
	trimmedScores := scores[1:7]
	//	平均分
	var sum int
	for _, score := range trimmedScores {
		sum += score
	}
	finalScore := float64(sum) / float64(len(trimmedScores))
	fmt.Println("平均分", finalScore)
	return finalScore, nil
}

// 找出最高分和最低分的评委
func findExtremeJudges(scores []int) (highestIndex, lowestIndex int) {
	maxScore := scores[0]
	minScore := scores[0]
	maxIndex := 0
	minIndex := 0
	for i, score := range scores {
		if score > maxScore {
			maxScore = score
			maxIndex = i
		}
		if score < minScore {
			minScore = score
			minIndex = i
		}
	}
	return maxIndex, minIndex
}

// 最佳评委和最差评委
// 使用了math.Abs函数来计算绝对值
func findBestAndWorstJudges(scores []int, finalScore float64) (bestIndex, worstIndex int) {
	var bestDiff, worstDiff float64
	bestIndex = 0
	worstIndex = 0
	for i, score := range scores {
		diff := float64(score) - finalScore
		if math.Abs(diff) < math.Abs(bestDiff) || bestDiff == 0 {
			bestDiff = diff
			bestIndex = i
		}
		if math.Abs(diff) > math.Abs(worstDiff) {
			worstDiff = diff
			worstIndex = i
		}
	}
	return bestIndex, worstIndex
}
func main() {

	scores := []int{9, 8, 10, 9, 7, 6, 5, 4}
	// 最终得分
	finalScore, err := calculateFinalScore(scores)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("最终得分: %.2f\n", finalScore)

	// 最高分和最低分的评委
	highestIndex, lowestIndex := findExtremeJudges(scores)
	fmt.Printf("最高分的评委: %d, 分数: %d\n", highestIndex, scores[highestIndex])
	fmt.Printf("最低分的评委: %d, 分数: %d\n", lowestIndex, scores[lowestIndex])

	// 最佳评委和最差评委
	bestIndex, worstIndex := findBestAndWorstJudges(scores, finalScore)
	fmt.Printf("最佳评委: %d, 分数: %d, 与最终得分相差: %.2f\n", bestIndex, scores[bestIndex], float64(scores[bestIndex])-finalScore)
	fmt.Printf("最差评委: %d, 分数: %d, 与最终得分相差: %.2f\n", worstIndex, scores[worstIndex], float64(scores[worstIndex])-finalScore)
}
