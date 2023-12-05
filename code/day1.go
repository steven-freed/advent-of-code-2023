package code

import (
	"fmt"
	"strings"
)

func find(haystack string, needle string) int {
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}

func rfind(haystack string, needle string) int {
	for i := len(haystack) - len(needle); i >= 0; i-- {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}

func Trebuchet_Part2(input string) int {
	numbers := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9}
	new_input := ""
	for _, line := range strings.Split(input, "\n") {
		new_line := line
		min_len, min_index := 0, len(line)
		for wrd, num := range numbers {
			if i := find(line, wrd); i < min_index && i > -1 {
				min_index, min_len = i, len(wrd)
			}
			if j := find(line, fmt.Sprintf("%d", num)); j < min_index && j > -1 {
				min_index, min_len = j, 1
			}
		}
		if _, ok := numbers[line[min_index:min_index+min_len]]; ok {
			new_line = fmt.Sprintf("%v%v%v", line[:min_index], numbers[line[min_index:min_index+min_len]], line[min_index+min_len:])
		}
		max_len, max_index := 0, -1
		for wrd, num := range numbers {
			if i := rfind(new_line, wrd); i > max_index && i > -1 {
				max_index, max_len = i, len(wrd)
			}
			if j := rfind(new_line, fmt.Sprintf("%d", num)); j > max_index && j > -1 {
				max_index, max_len = j, 1
			}
		}
		if _, ok := numbers[new_line[max_index:max_index+max_len]]; ok {
			new_line = fmt.Sprintf("%v%v%v", new_line[:max_index], numbers[new_line[max_index:max_index+max_len]], new_line[max_index+max_len:])
		}
		new_input += fmt.Sprintf("%v\n", new_line)
	}
	return Trebuchet_Part1(new_input[:len(new_input)-1])
}

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
