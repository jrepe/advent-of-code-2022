package main

import (
	"advent-of-code-2022/utils"
	"fmt"
	"strings"
)

const inputFilePath = "./day2/input.txt"
const testInputFilePath = "./day2/test_input.txt"

func main() {
	input := utils.ReadFileToStringAndSplit(inputFilePath, "\n")
	result1 := problem1(input)
	fmt.Println("RESULT 1: ", result1)

	result2 := problem2(input)
	fmt.Println("RESULT 2: ", result2)
}

func problem1(input []string) int {
	var result int
	for _, round := range input {
		options := strings.Split(round, " ")
		result += calculateRoundResult(options[0], options[1])
	}
	return result
}

func problem2(input []string) int {
	var result int
	for _, round := range input {
		options := strings.Split(round, " ")
		result += calculateRoundPrediction(options[0], options[1])
	}
	return result
}

// Rock: A X
// Paper: B Y
// Scissors: C Z
func calculateRoundResult(player1 string, player2 string) int {
	var outcome int
	if (player1 == "A" && player2 == "Y") ||
		(player1 == "B" && player2 == "Z") ||
		(player1 == "C" && player2 == "X") {
		outcome += 6
	} else if (player1 == "A" && player2 == "X") ||
		(player1 == "B" && player2 == "Y") ||
		(player1 == "C" && player2 == "Z") {
		outcome += 3
	}

	if player2 == "X" {
		outcome += 1
	} else if player2 == "Y" {
		outcome += 2
	} else {
		outcome += 3
	}

	return outcome
}

// X lose
// Y draw
// Z win
func calculateRoundPrediction(player1 string, gameOutcome string) int {
	var player2 string
	if player1 == "A" && gameOutcome == "X" {
		player2 = "Z"
	} else if player1 == "A" && gameOutcome == "Y" {
		player2 = "X"
	} else if player1 == "A" && gameOutcome == "Z" {
		player2 = "Y"
	} else if player1 == "B" && gameOutcome == "X" {
		player2 = "X"
	} else if player1 == "B" && gameOutcome == "Y" {
		player2 = "Y"
	} else if player1 == "B" && gameOutcome == "Z" {
		player2 = "Z"
	} else if player1 == "C" && gameOutcome == "X" {
		player2 = "Y"
	} else if player1 == "C" && gameOutcome == "Y" {
		player2 = "Z"
	} else if player1 == "C" && gameOutcome == "Z" {
		player2 = "X"
	}
	return calculateRoundResult(player1, player2)
}
