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

type cube struct {
	red   int
	green int
	blue  int
}
type games map[int][]cube

func CubeConundrum_Part1(input string, red int, green int, blue int) (int, error) {
	games := make(games, 0)
	for _, line := range strings.Split(input, "\n") {
		cubeSets := make([]cube, 0)
		gameNum, err := strconv.Atoi(strings.TrimSpace(line[strings.Index(line, GAME)+len(GAME) : strings.Index(line, COLON)]))
		if err != nil {
			return 0, errors.New("unable to convert game number to int")
		}
		cubeSetsStr := strings.Split(strings.TrimSpace(strings.Split(line, COLON)[1]), ";")
		for _, set := range cubeSetsStr {
			c := cube{}
			for _, color := range strings.Split(set, ",") {
				colorData := strings.Split(strings.TrimSpace(color), " ")
				num, err := strconv.Atoi(colorData[0])
				if err != nil {
					return 0, errors.New("unable to convert cube set number to int")
				}
				switch colorData[1] {
				case "red":
					c.red = num
				case "green":
					c.green = num
				case "blue":
					c.blue = num
				}
			}
			cubeSets = append(cubeSets, c)
		}
		games[gameNum] = cubeSets
	}
	playableSum := 0
	for gameNum, cubeSets := range games {
		gamePossible := true
		for _, cube := range cubeSets {
			if cube.red > red || cube.green > green || cube.blue > blue {
				gamePossible = false
			}
		}
		if gamePossible {
			playableSum += gameNum
		}
	}
	return playableSum, nil
}
