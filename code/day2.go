package code

import (
	"errors"
	"strconv"
	"strings"
)

const (
	GAME  = "Game"
	COLON = ":"
)

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
