package days

import (
	"advent-of-code-2024/utils"
	"fmt"
	"strconv"
	"strings"
)

func getInput() (map[int][]int, [][]int) {
	input := utils.GetInputLines(2024, 5)

	orderingRules := make(map[int][]int)
	var updates [][]int

	var parsingUpdates bool
	for _, line := range input {
		if line == "" {
			parsingUpdates = true
			continue
		}

		if !parsingUpdates {
			parts := strings.Split(line, "|")

			before, err := strconv.Atoi(parts[0])
			utils.Check(err)

			after, err := strconv.Atoi(parts[1])
			utils.Check(err)

			orderingRules[before] = append(orderingRules[before], after)
		} else {
			parts := strings.Split(line, ",")
			update := make([]int, len(parts))

			for i, part := range parts {
				num, err := strconv.Atoi(part)
				utils.Check(err)

				update[i] = num
			}
			updates = append(updates, update)
		}

	}

	return orderingRules, updates
}

func Day5Part1() int {
	orderingRules, updates := getInput()

	//fmt.Printf("%v\n", orderingRules)
	//fmt.Printf("%v\n", updates)

	for _, update := range updates {
		for i := 0; i < len(update)-1; i++ {
			num := update[i]
			next := update[i+1:]
			rules := orderingRules[num]
			fmt.Printf("%d\t%v\t|\t%v\n", num, rules, next)
		}
	}

	return -1
}
