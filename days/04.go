package days

import "advent-of-code-2024/utils"

const XMAS = "XMAS"
const XMAS_LEN = len(XMAS)

func check(lines []string, i int, j int, vStep int, hStep int) bool {
	iFinal := i + (vStep * (XMAS_LEN - 1))
	jFinal := j + (hStep * (XMAS_LEN - 1))

	if iFinal < 0 || iFinal >= len(lines) {
		return false
	}

	if jFinal < 0 || jFinal >= len(lines[0]) {
		return false
	}

	for mul, checkChar := range XMAS {
		line := lines[i+(vStep*mul)]
		char := rune(line[j+(hStep*mul)])
		if char != checkChar {
			return false
		}
	}

	return true
}

func Day4Part1() int {
	input := utils.GetInput(2024, 4)

	var count int
	for i, line := range input {
		for j, char := range []rune(line) {
			if char != 'X' {
				continue
			}

			for v := -1; v <= 1; v++ {
				for h := -1; h <= 1; h++ {
					if check(input, i, j, v, h) {
						count++
					}
				}
			}

		}
	}

	return count
}
