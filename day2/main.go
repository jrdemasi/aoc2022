package main

import (
	"fmt"
	"os"
	"strings"
)

func readInput() (input string) {
	data, err := os.ReadFile("input")
	if err != nil {
		fmt.Println(err)
	}
	input = string(data)
	return
}

func getPartOneScore(enc string) (value int) {
	moves := strings.Fields(enc)
	if moves[1] == "X" {
		value = 1
	} else if moves[1] == "Y" {
		value = 2
	} else {
		value = 3
	}
	switch moves[1] {
	case "X":
		if moves[0] == "A" {
			value = value + 3
		} else if moves[0] == "B" {
			value = value + 0
		} else {
			value = value + 6
		}
	case "Y":
		if moves[0] == "A" {
			value = value + 6
		} else if moves[0] == "B" {
			value = value + 3
		} else {
			value = value + 0
		}
	case "Z":
		if moves[0] == "A" {
			value = value + 0
		} else if moves[0] == "B" {
			value = value + 6
		} else {
			value = value + 3
		}
	}
	return
}

func getPartTwoScore(enc string) (value int) {
	moves := strings.Fields(enc)
	switch moves[1] {
	case "X":
		value = 0
		if moves[0] == "A" {
			value = value + 3
		} else if moves[0] == "B" {
			value = value + 1
		} else {
			value = value + 2
		}
	case "Y":
		value = 3
		if moves[0] == "A" {
			value = value + 1
		} else if moves[0] == "B" {
			value = value + 2
		} else {
			value = value + 3
		}
	case "Z":
		value = 6
		if moves[0] == "A" {
			value = value + 2
		} else if moves[0] == "B" {
			value = value + 3
		} else {
			value = value + 1
		}
	}
	return
}

func partOne(rawInput string) {
	split := strings.Split(rawInput, "\n")
	score := 0
	for i := 0; i < len(split); i++ {
		if strings.TrimSpace(split[i]) != "" {
			score = score + getPartOneScore(split[i])
		}
	}
	fmt.Printf("The answer to part one is: %v\n", score)
}

func partTwo(rawInput string) {
	split := strings.Split(rawInput, "\n")
	score := 0
	for i := 0; i < len(split); i++ {
		if strings.TrimSpace(split[i]) != "" {
			score = score + getPartTwoScore(split[i])
		}
	}
	fmt.Printf("The answer to part two is: %v\n", score)
}

func main() {
	rawInput := readInput()
	partOne(rawInput)
	partTwo(rawInput)
}
