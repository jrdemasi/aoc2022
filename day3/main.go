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

func stringValueOf(letter string) (value int) {
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for i := 0; i < len(letters); i++ {
		if letter == string(letters[i]) {
			value = i + 1
		}
	}
	return
}

func partTwo(input string) {
	split := strings.Split(input, "\n")
	value := 0
	for i := 0; i < len(split); i = i + 3 {
		for l := 0; l < len(split[i]); l++ {
			if strings.Contains(split[i+1], string(split[i][l])) && strings.Contains(split[i+2], string(split[i][l])) {
				value = value + stringValueOf(string(split[i][l]))
				break
			}
		}
	}
	fmt.Printf("The answer to part two is %v\n", value)
}

func partOne(input string) {
	split := strings.Split(input, "\n")
	value := 0
	for i := 0; i < len(split); i++ {
		if strings.TrimSpace(split[i]) != "" {
			var ruckOne string
			var ruckTwo string
			numItems := len(split[i]) / 2
			ruckOne = split[i][0:numItems]
			ruckTwo = split[i][numItems:]
			for l := 0; l < len(ruckOne); l++ {
				if strings.Contains(ruckTwo, string(ruckOne[l])) {
					value = value + stringValueOf(string(ruckOne[l]))
					break
				}
			}
		}
	}
	fmt.Printf("The answer to part one is: %v\n", value)
}

func main() {
	rawInput := readInput()
	partOne(rawInput)
	partTwo(rawInput)
}
