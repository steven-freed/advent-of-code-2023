package code

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

type AlmanMap struct {
	source      string
	dest        string
	destCol     []int
	srcCol      []int
	rangeLenCol []int
}

func (m AlmanMap) srcToDest(src int) int {
	for row, rangeVal := range m.rangeLenCol {
		srcWithinRange := src >= m.srcCol[row] && src <= (m.srcCol[row]+rangeVal-1)
		if srcWithinRange {
			return m.destCol[row] + (src - m.srcCol[row])
		}
	}
	return src
}

func almanacToAlmanMaps(input string) (map[string]AlmanMap, error) {
	lines := strings.Split(input, "\n")
	maps := make(map[string]AlmanMap, 0)
	for i := 0; i < len(lines); i++ {
		if len(lines[i]) > 0 && strings.Contains(lines[i], "map") {
			// generate almanac maps
			mapName := strings.Split(lines[i][:strings.Index(lines[i], " ")], "-")
			m := &AlmanMap{
				source:      mapName[0],
				dest:        mapName[2],
				destCol:     make([]int, 0),
				srcCol:      make([]int, 0),
				rangeLenCol: make([]int, 0),
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
				m.destCol = append(m.destCol, dest)
				m.srcCol = append(m.srcCol, src)
				m.rangeLenCol = append(m.rangeLenCol, rangeLen)
				i++ // skip over blank line after map table
			}
			// add src map
			maps[mapName[0]] = *m
		}
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
	minLocation := math.MaxFloat64
	for _, seed := range seeds {
		if key == "seed" {
			nextSrc = seed
		}
		for key != "location" {
			nextSrc = almanMaps[key].srcToDest(nextSrc)
			key = almanMaps[key].dest
		}
		minLocation = math.Min(minLocation, float64(nextSrc))
		key = "seed"
	}
	return int(minLocation), nil
}
