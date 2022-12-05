package main

import (
	"advent-of-code-2022/utils"
	"fmt"
	"strconv"
	"strings"
)

const inputFilePath = "./day5/input.txt"
const testInputFilePath = "./day5/test_input.txt"

type dequeue []string

func (q dequeue) enqueue(item string) dequeue {
	return append(dequeue{item}, q...)
}

func (q dequeue) dequeue() (dequeue, string) {
	l := len(q)
	return q[:l-1], q[l-1]
}

func (q dequeue) push(item string) dequeue {
	return append(q, dequeue{item}...)
}

func (q dequeue) pushAll(item dequeue) dequeue {
	return append(q, item...)
}

func (q dequeue) peek() string {
	return q[len(q)-1]
}

func main() {
	input := utils.ReadFileToStringAndSplit(inputFilePath, "\n\n")

	result1 := problem1(input)
	fmt.Println("RESULT 1: ", result1)

	result2 := problem2(input)
	fmt.Println("RESULT 2: ", result2)
}

func problem1(input []string) string {
	state := getInitialState(input[0])
	instr := parseInstructions(input[1])
	for _, move := range instr {
		state = shift(state, move)
	}
	elements := getResultingString(state)
	return strings.Join(elements, "")
}

func problem2(input []string) string {
	state := getInitialState(input[0])
	instr := parseInstructions(input[1])
	for _, move := range instr {
		state = shiftInOrder(state, move)
	}
	elements := getResultingString(state)
	return strings.Join(elements, "")
}

func shift(state map[int]dequeue, move []int) map[int]dequeue {
	var el string
	for i := 0; i < move[0]; i++ {
		state[move[1]], el = state[move[1]].dequeue()
		state[move[2]] = state[move[2]].push(el)
	}
	return state
}

func shiftInOrder(state map[int]dequeue, move []int) map[int]dequeue {
	var el string
	var interim dequeue
	for i := 0; i < move[0]; i++ {
		state[move[1]], el = state[move[1]].dequeue()
		interim = interim.enqueue(el)
	}
	state[move[2]] = state[move[2]].pushAll(interim)
	return state
}

func parseInstructions(input string) [][]int {
	instructions := make([][]int, 0)
	for _, line := range strings.Split(input, "\n") {
		var s []int
		for i, el := range strings.Fields(line) {
			if i%2 == 0 {
				continue
			}
			element, _ := strconv.Atoi(el)
			s = append(s, element)
		}
		instructions = append(instructions, s)
		s = nil
	}
	return instructions
}

func getInitialState(input string) map[int]dequeue {
	state := make(map[int]dequeue, 0)
	lines := strings.Split(input, "\n")
	for _, line := range lines[:len(lines)-1] {
		for i, stack := range strings.Split(line, " ") {
			if stack == "x" {
				continue
			}
			state[i+1] = state[i+1].enqueue(stack)
		}
	}
	return state
}

func getResultingString(state map[int]dequeue) []string {
	i := 1
	var elements []string
	for {
		if v, ok := state[i]; ok {
			elements = append(elements, v.peek())
		} else {
			break
		}
		i++
	}
	return elements
}
