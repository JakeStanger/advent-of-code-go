package days

import (
	"advent-of-code-2024/utils"
	"strconv"
	"strings"
)

func Day2Part1() int {
	input := utils.GetInput(2024, 2)

	var numSafe int

	for _, line := range input {
		lineValues := strings.Split(line, " ")

		isSafe := true

		prev := 0
		var ascending bool

		for i, value := range lineValues {
			valueNum, err := strconv.Atoi(value)
			utils.Check(err)

			if i == 0 {
				prev = valueNum
				continue
			}

			diff := valueNum - prev

			if diff == 0 || diff > 3 || diff < -3 {
				isSafe = false
				break
			}

			currAscending := diff > 0

			if i > 1 && currAscending != ascending {
				isSafe = false
				break
			}

			prev = valueNum
			ascending = currAscending
		}

		if isSafe {
			numSafe++
		}
	}

	return numSafe
}
