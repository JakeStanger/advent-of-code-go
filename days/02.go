package days

import (
	"advent-of-code-2024/utils"
	"strconv"
	"strings"
)

func getUnsafeIndex(lineValues []int, ignoreIndex int) int {
	for i := 1; i < len(lineValues); i++ {
		if i == ignoreIndex {
			continue
		}

		num := lineValues[i]
		prev := lineValues[i-1]

		if num == prev {
			return i
		}

		diff := num - prev
		if diff > 3 || diff < -3 {
			return i
		}

		ascending := diff > 0
		if i > 1 {
			prevDiff := lineValues[i-1] - lineValues[i-2]
			prevAscending := prevDiff > 0

			if ascending != prevAscending {
				return i
			}
		}
	}

	return -1
}

func Day2Part1() int {
	input := utils.GetInputLines(2024, 2)

	var numSafe int

	for _, line := range input {
		lineValuesStr := strings.Split(line, " ")
		lineValues := make([]int, len(lineValuesStr))

		for i, value := range lineValuesStr {
			num, err := strconv.Atoi(value)
			utils.Check(err)
			lineValues[i] = num
		}

		if getUnsafeIndex(lineValues, -1) == -1 {
			numSafe++
		}
	}

	return numSafe
}

func Day2Part2() int {
	input := utils.GetInputLines(2024, 2)

	var numSafe int

	for _, line := range input {
		lineValuesStr := strings.Split(line, " ")
		lineValues := make([]int, len(lineValuesStr))

		for i, value := range lineValuesStr {
			num, err := strconv.Atoi(value)
			utils.Check(err)
			lineValues[i] = num
		}

		unsafeIndex := getUnsafeIndex(lineValues, -1)
		safe := unsafeIndex == -1
		if !safe {
			//for i := range lineValues {
			//	idx := getUnsafeIndex(lineValues, i)
			//	if idx == -1 {
			//		safe = true
			//		break
			//	}
			//}

			//slice2 := lineValues[1:]
			//
			////println(unsafeIndex)
			// safe = getUnsafeIndex(lineValues, unsafeIndex) == -1 || getUnsafeIndex(lineValues, unsafeIndex-1) == -1

			for i := range lineValues {
				cpy := utils.Copy(lineValues)
				cpy = utils.Remove(cpy, i)

				idx := getUnsafeIndex(cpy, -1)
				if idx == -1 {
					safe = true
					break
				}
			}
		}

		if safe {
			numSafe++
		}
	}

	return numSafe
}
