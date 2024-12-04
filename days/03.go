package days

import (
	"advent-of-code-2024/utils"
	"regexp"
	"strconv"
)

func Day3Part1() int {
	input := utils.GetInput(2024, 3)

	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	var total int

	matches := re.FindAllStringSubmatch(input, -1)
	for _, match := range matches {
		n1, err := strconv.Atoi(match[1])
		utils.Check(err)
		n2, err := strconv.Atoi(match[2])
		utils.Check(err)

		total += n1 * n2
	}

	return total

}
