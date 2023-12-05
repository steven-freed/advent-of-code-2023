package main

import (
	"advent-of-code-2023/code"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func run_day(day int, input string) {
	switch day {
	case 1:
		fmt.Println("Day 1: Trebuchet!?")
		sum := code.Trebuchet_Part1(input)
		fmt.Printf("Part 1 Answer: %d\n", sum)
		sum = code.Trebuchet_Part2(input)
		fmt.Printf("Part 2 Answer: %d\n", sum)
	}
}

func main() {
	days := []int{1}
	for _, day := range days {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Enter input: ")
		input := ""
		for scanner.Scan() {
			input += fmt.Sprintf("%s\n", strings.TrimSpace(scanner.Text()))
		}
		input = input[:len(input)-1]
		run_day(day, input)
	}
}
