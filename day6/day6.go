package main

import (
	"advent-of-code-2022/utils"
	"fmt"
)

const inputFilePath = "./day6/input.txt"
const testInputFilePath = "./day6/test_input.txt"

func main() {
	input := utils.ReadFileToStringAndSplit(inputFilePath, "")

	result1 := problem1(input)
	fmt.Println("RESULT 1: ", result1)

	result2 := problem2(input)
	fmt.Println("RESULT 2: ", result2)
}

func problem1(input []string) int {
	return solveWithWindowSize(input, 4)
}

func problem2(input []string) int {
	return solveWithWindowSize(input, 14)
}

func solveWithWindowSize(input []string, size int) int {
	for i := size; i < len(input); i++ {
		if isUniqueCharacters(input[i-size : i]) {
			return i
		}
	}
	return -1
}

func isUniqueCharacters(input []string) bool {
	m := make(map[string]struct{})
	for _, v := range input {
		if _, ok := m[v]; ok {
			return false
		}
		m[v] = struct{}{}
	}
	return true
}
