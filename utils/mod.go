package utils

import (
	"fmt"
	"os"
	"strings"
)

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func GetInput(year int, day int) []string {
	fileName := fmt.Sprintf("inputs/%d/%02d", year, day)
	bytes, err := os.ReadFile(fileName)
	Check(err)

	input := string(bytes)
	lines := strings.Split(input, "\n")

	return lines
}

func Copy[T any](slice []T) []T {
	tmp := make([]T, len(slice))
	copy(tmp, slice)
	return tmp
}

func Remove[T any](slice []T, index int) []T {
	return append(slice[:index], slice[index+1:]...)
}
