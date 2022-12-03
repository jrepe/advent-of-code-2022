package main

import (
	"advent-of-code-2022/utils"
	"fmt"
	"sort"
	"strings"
)

const inputFilePath = "./day1/input.txt"
const testInputFilePath = "./day1/test_input.txt"

func main() {
	input := utils.ReadFileToStringAndSplit(inputFilePath, "\n\n")
	result1 := problem1(input)
	fmt.Println("RESULT 1: ", result1)

	result2 := problem2(input)
	fmt.Println("RESULT 2: ", result2)
}

func problem1(input []string) int {
	max := 0
	for _, line := range input {
		items := utils.ConvertStringSliceToInt(strings.Split(line, "\n"))
		elfSum := utils.SumOfIntItems(items)
		if elfSum >= max {
			max = elfSum
		}
	}
	return max
}

func problem2(input []string) int {
	max := make([]int, 3)
	for _, line := range input {
		items := utils.ConvertStringSliceToInt(strings.Split(line, "\n"))
		elfSum := utils.SumOfIntItems(items)
		max = updateLargestSumSlice(max, elfSum)
	}
	return utils.SumOfIntItems(max)
}

func updateLargestSumSlice(max []int, candidate int) []int {
	sort.Ints(max)
	for i, r := range max {
		if candidate > r {
			max[i] = candidate
			break
		}
	}
	return max
}
