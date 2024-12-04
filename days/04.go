package days

import "advent-of-code-2024/utils"

const XMAS = "XMAS"

func check(lines []string, i int, j int, vStep int, hStep int, word string) bool {
	wordLen := len(word)

	iFinal := i + (vStep * (wordLen - 1))
	jFinal := j + (hStep * (wordLen - 1))

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
	input := utils.GetInputLines(2024, 4)

	var count int
	for i, line := range input {
		for j, char := range []rune(line) {
			if char != 'X' {
				continue
			}

			for v := -1; v <= 1; v++ {
				for h := -1; h <= 1; h++ {
					if check(input, i, j, v, h, XMAS) {
						count++
					}
				}
			}

		}
	}

	return count
}

func Day4Part2() int {
	input := utils.GetInputLines(2024, 4)

	var count int
	for i := 1; i < len(input)-1; i++ {
		for j := 1; j < len(input[i])-1; j++ {
			if j == 0 || input[i][j] != 'A' {
				continue
			}

			topLeft := rune(input[i-1][j-1])
			topRight := rune(input[i-1][j+1])
			bottomLeft := rune(input[i+1][j-1])
			bottomRight := rune(input[i+1][j+1])

			check1 := (topLeft == 'M' && bottomRight == 'S') || topLeft == 'S' && bottomRight == 'M'
			check2 := (topRight == 'M' && bottomLeft == 'S') || topRight == 'S' && bottomLeft == 'M'

			if check1 && check2 {
				count++
			}
		}
	}

	return count
}
