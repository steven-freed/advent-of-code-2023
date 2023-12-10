package code

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type AlmanMap struct {
	source           string
	dest             string
	sourceRangeStart []int
	destRangeStart   []int
	rangeLen         []int
}

func (m AlmanMap) String() string {
	return fmt.Sprintf("{source: %v, dest: %v, sourceRangeStart: %v, destRangeStart: %v, rangeLen: %v}", m.source, m.dest, m.sourceRangeStart, m.destRangeStart, m.rangeLen)
}

func NewAlmanMap(source string, dest string) *AlmanMap {
	a := AlmanMap{
		source: source,
		dest:   dest,
	}
	a.sourceRangeStart = make([]int, 0)
	a.destRangeStart = make([]int, 0)
	a.rangeLen = make([]int, 0)
	return &a
}

func SeedFertilizer_Part1(input string) (int, error) {
	lines := strings.Split(input, "\n")
	seeds := make([]int, 0)
	maps := make(map[string]AlmanMap, 0)
	for i := 0; i < len(lines); i++ {
		if len(lines[i]) == 0 {
			continue
		} else if !strings.Contains(lines[i], "map") {
			// generate seeds slice
			sliceaSeeds := strings.Split(strings.TrimSpace(lines[i][strings.Index(lines[i], ":")+1:]), " ")
			for _, s := range sliceaSeeds {
				s, err := strconv.Atoi(s)
				if err != nil {
					return 0, errors.New("unable to parse seeds")
				}
				seeds = append(seeds, s)
			}
		} else {
			// generate almanac maps
			mapName := strings.Split(lines[i][:strings.Index(lines[i], " ")], "-")
			m := NewAlmanMap(mapName[0], mapName[2])
			i++
			for i < len(lines) && len(lines[i]) > 0 {
				nums := strings.Split(lines[i], " ")
				for i, n := range nums {
					n, err := strconv.Atoi(n)
					if err != nil {
						return 0, errors.New("unable to parse map")
					}
					switch i {
					case 0:
						m.sourceRangeStart = append(m.sourceRangeStart, n)
					case 1:
						m.destRangeStart = append(m.destRangeStart, n)
					case 2:
						m.rangeLen = append(m.rangeLen, n)
					}
				}
				i++
			}
			maps[strings.Join(mapName, "")] = *m
		}
	}
	return 0, nil
}
