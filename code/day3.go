package code

import (
	"errors"
	"strconv"
	"strings"
)

func schematicToGraph(schematic string) [][]rune {
	rows := strings.Split(schematic, "\n")
	rowCount := len(rows)
	var colCount int
	if len(rows) > 0 {
		colCount = len(rows[0])
	}
	schemGraph := make([][]rune, rowCount)
	for i, row := range rows {
		schemGraphRow := make([]rune, colCount)
		for j, col := range row {
			schemGraphRow[j] = col
		}
		schemGraph[i] = schemGraphRow
	}
	return schemGraph
}

func isValidPartNumber(r int, c int, graph [][]rune, rows int, cols int) bool {
	checks := [][]int{
		{0, 1},
		{1, 0},
		{-1, 0},
		{0, -1},
		{1, 1},
		{-1, 1},
		{1, -1},
		{-1, -1},
	}
	for _, el := range checks {
		i, j := r+el[0], c+el[1]
		if i < rows && i > -1 && j < cols && j > -1 && (graph[i][j] < 48 || graph[i][j] > 57) && graph[i][j] != 46 {
			return true
		}
	}
	return false
}

type coord [2]int
type gearMeta struct {
	ratio       int
	adjPartNums int
}

func findGears(r int, c int, graph [][]rune, rows int, cols int) map[coord]bool {
	checks := [][]int{
		{0, 1},
		{1, 0},
		{-1, 0},
		{0, -1},
		{1, 1},
		{-1, 1},
		{1, -1},
		{-1, -1},
	}
	gears := make(map[coord]bool)
	for _, el := range checks {
		i, j := r+el[0], c+el[1]
		if i < rows && i > -1 && j < cols && j > -1 && graph[i][j] == 42 {
			gears[coord{i, j}] = true
		}
	}
	return gears
}

func GearRatios_Part2(input string) (int, error) {
	// convert input to graph
	schemGraph := schematicToGraph(input)
	rowCount, colCount := len(schemGraph), len(schemGraph[0])
	gearRatios := make(map[coord]gearMeta)
	for r, _ := range schemGraph {
		startCol := -1
		gearsForDigit := make(map[coord]bool)
		for c, _ := range schemGraph[r] {
			digit := isDigit(schemGraph[r][c])
			if digit {
				if startCol < 0 {
					startCol = c
					gearsForDigit = make(map[coord]bool)
				}
				gears := findGears(r, c, schemGraph, rowCount, colCount)
				for k, v := range gears {
					gearsForDigit[k] = v
				}
			}
			if startCol > -1 && (!digit || c == rowCount-1) {
				if (c == rowCount-1) && digit {
					c = rowCount
				}
				partNum, err := strconv.Atoi(string(schemGraph[r][startCol:c]))
				if err != nil {
					return 0, errors.New("unable to convert part number to int")
				}
				for k, _ := range gearsForDigit {
					gearRatioCopy, ok := gearRatios[k]
					if !ok {
						gearRatios[k] = gearMeta{ratio: partNum, adjPartNums: 1}
					} else {
						gearRatioCopy.ratio *= partNum
						gearRatioCopy.adjPartNums += 1
						gearRatios[k] = gearRatioCopy

					}
				}
			}
			if !digit {
				startCol = -1
			}
		}
	}
	gearRatiosSum := 0
	for _, v := range gearRatios {
		if v.adjPartNums == 2 {
			gearRatiosSum += v.ratio
		}
	}
	return gearRatiosSum, nil
}

func GearRatios_Part1(input string) (int, error) {
	// convert input to graph
	schemGraph := schematicToGraph(input)
	rowCount, colCount := len(schemGraph), len(schemGraph[0])
	// sum valid part numbers
	partNumSum := 0
	for r, _ := range schemGraph {
		startCol := -1
		validPartNum := false
		for c, _ := range schemGraph[r] {
			digit := isDigit(schemGraph[r][c])
			if digit {
				if startCol < 0 {
					startCol = c
				}
				if !validPartNum && isValidPartNumber(r, c, schemGraph, rowCount, colCount) {
					validPartNum = true
				}
			}
			if (startCol > -1 && validPartNum) && (!digit || c == rowCount-1) {
				if (c == rowCount-1) && digit {
					c = rowCount
				}
				partNum, err := strconv.Atoi(string(schemGraph[r][startCol:c]))
				if err != nil {
					return 0, errors.New("unable to convert part number to int")
				}
				partNumSum += partNum
			}
			if !digit {
				validPartNum = false
				startCol = -1
			}
		}
	}
	return partNumSum, nil
}
