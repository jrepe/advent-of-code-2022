package main

import (
	"advent-of-code-2022/utils"
	"fmt"
	"unicode"
)

const inputFilePath = "./day3/input.txt"
const testInputFilePath = "./day3/test_input.txt"

func main() {
	input := utils.ReadFileToStringAndSplit(inputFilePath, "\n")
	result1 := problem1(input)
	fmt.Println("RESULT 1: ", result1)

	result2 := problem2(input)
	fmt.Println("RESULT 2: ", result2)
}

func problem1(input []string) int {
	var duplicates []rune
	for _, line := range input {
		duplicates = append(duplicates, getDuplicateItems(line)...)
	}
	priorities := getItemPriorities(duplicates)
	return utils.SumOfIntItems(priorities)
}

func problem2(input []string) int {
	var duplicates []rune
	var groupLines []string
	for _, line := range input {
		groupLines = append(groupLines, line)
		if len(groupLines) == 3 {
			duplicates = append(duplicates, getDuplicateItemsBetweenGroups(groupLines)...)
			groupLines = nil
		}
	}
	priorities := getItemPriorities(duplicates)
	return utils.SumOfIntItems(priorities)
}

func getDuplicateItems(input string) []rune {
	first := utils.SplitStringByIndex(input, 0, len(input)/2)
	second := utils.SplitStringByIndex(input, len(input)/2, len(input))
	return utils.IntersectBetweenStrings(first, second)
}

func getItemPriorities(items []rune) (priorities []int) {
	for _, item := range items {
		priorities = append(priorities, getPriority(item))
	}
	return
}

func getPriority(item rune) int {
	if unicode.IsUpper(item) {
		return int(item) - 38
	}
	return int(item) - 96
}

func getDuplicateItemsBetweenGroups(input []string) []rune {
	return utils.IntersectBetweenStrings(
		string(utils.IntersectBetweenStrings(input[0], input[1])),
		input[2])
}
