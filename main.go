package main

import (
	"advent-of-code-2023/code"
	"bufio"
	"fmt"
	"log"
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
	case 2:
		fmt.Println("Day 2: Cube Conundrum")
		sum, err := code.CubeConundrum_Part1(input, 12, 13, 14)
		if err != nil {
			fmt.Printf("error: %d\n", err)
		}
		fmt.Printf("Part 1 Answer: %d\n", sum)
		sum, err = code.CubeConundrum_Part2(input)
		if err != nil {
			fmt.Printf("error: %d\n", err)
		}
		fmt.Printf("Part 2 Answer: %d\n", sum)
	case 3:
		fmt.Println("Day 3: Gear Ratios")
		sum, err := code.GearRatios_Part1(input)
		if err != nil {
			fmt.Printf("error: %d\n", err)
		}
		fmt.Printf("Part 1 Answer: %d\n", sum)
		sum, err = code.GearRatios_Part2(input)
		if err != nil {
			fmt.Printf("error: %d\n", err)
		}
		fmt.Printf("Part 2 Answer: %d\n", sum)
	case 4:
		fmt.Println("Day 4: Scratchcards")
		sum, err := code.Scratchcards_Part1(input)
		if err != nil {
			fmt.Printf("error: %d\n", err)
		}
		fmt.Printf("Part 1 Answer: %d\n", sum)
		sum, err = code.Scratchcards_Part2(input)
		if err != nil {
			fmt.Printf("error: %d\n", err)
		}
		fmt.Printf("Part 2 Answer: %d\n", sum)
	case 5:
		fmt.Println("Day 5: If You Give A Seed A Fertilizer")
		sum, err := code.SeedFertilizer_Part1(input)
		if err != nil {
			fmt.Printf("error: %d\n", err)
		}
		fmt.Printf("Part 1 Answer: %d\n", sum)
		sum, err = code.SeedFertilizer_Part2(input)
		if err != nil {
			fmt.Printf("error: %d\n", err)
		}
		fmt.Printf("Part 2 Answer: %d\n", sum)
	}
}

func main() {
	days := []int{5}
	for _, day := range days {
		file, err := os.Open(fmt.Sprintf("day%v_input.txt", day))
		if err != nil {
			log.Fatal("unable to read input file")
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		input := ""
		for scanner.Scan() {
			input += fmt.Sprintf("%s\n", strings.TrimSpace(scanner.Text()))
		}
		input = input[:len(input)-1]
		run_day(day, input)
	}
}
