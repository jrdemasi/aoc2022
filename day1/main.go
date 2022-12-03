package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
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

/* Takes the input and finds the greatest calorie count among the elves */
func part_one(input string) {
	split := strings.Split(input, "\n")
	max := 0
	elf := 0
	for i := 0; i < len(split); i++ {
		if strings.TrimSpace(split[i]) == "" {
			elf = 0
		} else {
			cal, _ := strconv.Atoi(split[i])
			elf = elf + cal
			if elf > max {
				max = elf
			}
		}
	}
	fmt.Println(max)
}

/* Takes the input and finds the sum of the calories of the top 3 elves */
func part_two(input string) {
	split := strings.Split(input, "\n")
	elf := 0
	var elves []int
	for i := 0; i < len(split); i++ {
		if strings.TrimSpace(split[i]) == "" {
			elves = append(elves, elf)
			elf = 0
		} else {
			cal, _ := strconv.Atoi(split[i])
			elf = elf + cal
		}
	}
	sort.Slice(elves, func(i, j int) bool {
		return elves[i] > elves[j]
	})
	sum := elves[0] + elves[1] + elves[2]
	fmt.Println(sum)
}

func main() {
	raw_input := readInput()
	part_one(raw_input)
	part_two(raw_input)
}
