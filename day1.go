package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Trebuchet_Part1(input string) int {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		num := 0
		for _, n := range line {
			if n >= 48 && n <= 57 {
				num += int(n) - 48
				break
			}
		}
		for i, _ := range line {
			n := line[len(line)-i-1]
			if n >= 48 && n <= 57 {
				if num > 0 {
					num *= 10
				}
				num += int(n) - 48
				break
			}
		}
		sum += num
	}
	return sum
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter input: ")
	input := ""
	for scanner.Scan() {
		input += fmt.Sprintf("%s\n", strings.TrimSpace(scanner.Text()))
	}
	sum := Trebuchet_Part1(input)
	fmt.Println("Day 1: Trebuchet!?")
	fmt.Println(fmt.Sprintf("Answer: %d", sum))
}
