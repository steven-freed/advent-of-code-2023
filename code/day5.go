package code

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type AlmanMap struct {
	source    string
	dest      string
	srcToDest map[int]int
}

func almanacToAlmanMaps(input string) (map[string]AlmanMap, error) {
	lines := strings.Split(input, "\n")
	maps := make(map[string]AlmanMap, 0)
	for i := 0; i < len(lines); i++ {
		if len(lines[i]) > 0 && strings.Contains(lines[i], "map") {
			// generate almanac maps
			mapName := strings.Split(lines[i][:strings.Index(lines[i], " ")], "-")
			m := &AlmanMap{
				source:    mapName[0],
				dest:      mapName[2],
				srcToDest: make(map[int]int),
			}
			i++ // skip over map name line
			// parse map table
			for i < len(lines) && len(lines[i]) > 0 {
				nums := strings.Split(lines[i], " ")
				dest, derr := strconv.Atoi(nums[0])
				src, serr := strconv.Atoi(nums[1])
				rangeLen, rerr := strconv.Atoi(nums[2])
				if derr != nil || serr != nil || rerr != nil {
					return nil, errors.New("unable to parse map")
				}
				for j := 0; j < rangeLen; j++ {
					srcNum := src + j
					destNum := dest + j
					m.srcToDest[srcNum] = destNum
				}
				i++ // skip over blank line after map table
			}
			// add src map
			maps[mapName[0]] = *m
		}
		fmt.Println("parsed line:", i+1, "of", len(lines))
	}
	return maps, nil
}

func SeedFertilizer_Part1(input string) (int, error) {
	seeds := make([]int, 0)
	for _, line := range strings.Split(input, "\n") {
		if len(line) > 0 && !strings.Contains(line, "map") {
			// generate seeds slice
			sliceaSeeds := strings.Split(strings.TrimSpace(line[strings.Index(line, ":")+1:]), " ")
			for _, s := range sliceaSeeds {
				s, err := strconv.Atoi(s)
				if err != nil {
					return 0, errors.New("unable to parse seeds")
				}
				seeds = append(seeds, s)
			}
			break
		}
	}
	almanMaps, err := almanacToAlmanMaps(input)
	if err != nil {
		return 0, err
	}
	key := "seed"
	nextSrc := 0
	minLocation := math.MaxInt64
	for _, seed := range seeds {
		if key == "seed" {
			nextSrc = seed
		}
		for key != "location" {
			if _, ok := almanMaps[key].srcToDest[nextSrc]; ok {
				nextSrc = almanMaps[key].srcToDest[nextSrc]
			}
			key = almanMaps[key].dest
		}
		minLocation = int(math.Min(float64(minLocation), float64(nextSrc)))
		key = "seed"
	}
	return minLocation, nil
}
