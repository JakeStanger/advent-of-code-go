package days

import (
	"advent-of-code-2024/utils"
	"math"
	"sort"
	"strconv"
	"strings"
)

func getLists() ([]int, []int) {
	lines := utils.GetInputLines(2024, 01)

	var left []int
	var right []int
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		parts := strings.Split(line, "   ")

		leftNum, err := strconv.Atoi(parts[0])
		utils.Check(err)

		rightNum, err := strconv.Atoi(parts[1])
		utils.Check(err)

		left = append(left, leftNum)
		right = append(right, rightNum)
	}

	return left, right
}

func Day1Part1() int {
	left, right := getLists()

	sort.Sort(sort.IntSlice(left))
	sort.Sort(sort.IntSlice(right))

	totalDistance := 0

	for i, leftNum := range left {
		rightNum := right[i]

		distance := math.Abs(float64(leftNum - rightNum))
		totalDistance += int(distance)
	}

	return totalDistance
}

func countOccurrences(nums []int, match int) int {
	matches := 0
	for _, num := range nums {
		if num == match {
			matches++
		}
	}

	return matches
}

func Day1Part2() int {
	left, right := getLists()

	totalSimilarity := 0

	for _, leftNum := range left {
		similarity := leftNum * countOccurrences(right, leftNum)
		totalSimilarity += similarity
	}

	return totalSimilarity
}
