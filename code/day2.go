package code

import (
	"errors"
	"strconv"
	"strings"
)

func CubeConundrum_Part2(input string) (int, error) {
	cubePowerSum := 0
	for _, line := range strings.Split(input, "\n") {
		cubeSetsStr := strings.Split(strings.TrimSpace(strings.Split(line, COLON)[1]), ";")
		colorMaxs := []int{0, 0, 0} // RGB ordering
		for _, set := range cubeSetsStr {
			for _, color := range strings.Split(set, ",") {
				colorData := strings.Split(strings.TrimSpace(color), " ")
				num, err := strconv.Atoi(colorData[0])
				if err != nil {
					return 0, errors.New("unable to convert cube set number to int")
				}
				switch {
				case colorData[1] == "red" && num > colorMaxs[0]:
					colorMaxs[0] = num
				case colorData[1] == "green" && num > colorMaxs[1]:
					colorMaxs[1] = num
				case colorData[1] == "blue" && num > colorMaxs[2]:
					colorMaxs[2] = num
				}
			}
		}
		for _, max := range colorMaxs[1:] {
			colorMaxs[0] *= max
		}
		cubePowerSum += colorMaxs[0]
	}
	return cubePowerSum, nil
}

func CubeConundrum_Part1(input string, red int, green int, blue int) (int, error) {
	playableSum := 0
	for _, line := range strings.Split(input, "\n") {
		gameNum, err := strconv.Atoi(strings.TrimSpace(line[strings.Index(line, GAME)+len(GAME) : strings.Index(line, COLON)]))
		if err != nil {
			return 0, errors.New("unable to convert game number to int")
		}
		gamePossible := true
		cubeSetsStr := strings.Split(strings.TrimSpace(strings.Split(line, COLON)[1]), ";")
		for _, set := range cubeSetsStr {
			for _, color := range strings.Split(set, ",") {
				colorData := strings.Split(strings.TrimSpace(color), " ")
				num, err := strconv.Atoi(colorData[0])
				if err != nil {
					return 0, errors.New("unable to convert cube set number to int")
				}
				if (colorData[1] == "red" && num > red) || (colorData[1] == "green" && num > green) || colorData[1] == "blue" && num > blue {
					gamePossible = false
				}
			}
		}
		if gamePossible {
			playableSum += gameNum
		}
	}
	return playableSum, nil
}
