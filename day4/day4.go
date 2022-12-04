package main

import (
	"advent-of-code-2022/utils"
	"fmt"
	"strconv"
	"strings"
)

const inputFilePath = "./day4/input.txt"
const testInputFilePath = "./day4/test_input.txt"

type section struct {
	start int
	end   int
}

func createSection(pair string) section {
	coord := strings.Split(pair, "-")
	first, _ := strconv.Atoi(coord[0])
	second, _ := strconv.Atoi(coord[1])
	return section{start: first, end: second}
}

func main() {
	input := utils.ReadFileToStringAndSplit(inputFilePath, "\n")
	result1 := problem1(input)
	fmt.Println("RESULT 1: ", result1)

	result2 := problem2(input)
	fmt.Println("RESULT 2: ", result2)
}

func problem1(input []string) int {
	var contained int
	for _, line := range input {
		pairs := strings.Split(line, ",")
		first := createSection(pairs[0])
		second := createSection(pairs[1])
		if isWithinRange(first, second) {
			contained += 1
		}
	}
	return contained
}

func problem2(input []string) int {
	var overlap int
	for _, line := range input {
		pairs := strings.Split(line, ",")
		first := createSection(pairs[0])
		second := createSection(pairs[1])
		if isOverlap(first, second) {
			overlap += 1
		}
	}
	return overlap
}

func isOverlap(first section, second section) bool {
	return first.start <= second.end && second.start <= first.end
}

func isWithinRange(first section, second section) bool {
	return (first.start <= second.start && first.end >= second.end) ||
		(second.start <= first.start && second.end >= first.end)
}
