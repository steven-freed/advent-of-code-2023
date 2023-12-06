package code

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func prettyPrint(graph [][]rune) {
	for i, _ := range graph {
		for j, _ := range graph[i] {
			fmt.Print(string(graph[i][j]))
		}
		fmt.Println()
	}
}

func isDigit(input rune) bool {
	return input >= 48 && input <= 57
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

func GearRatios_Part1(input string) (int, error) {
	// convert input to graph
	rows := strings.Split(input, "\n")
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
	// sum valid part numbers
	partNumSum := 0
	for r, _ := range schemGraph {
		startCol := -1
		validPartNum := false
		for c, _ := range schemGraph[r] {
			if isDigit(schemGraph[r][c]) {
				if startCol < 0 {
					startCol = c
				}
				if !validPartNum && isValidPartNumber(r, c, schemGraph, rowCount, colCount) {
					validPartNum = true
				}
				continue
			}
			if startCol > -1 && validPartNum {
				partNum, err := strconv.Atoi(string(schemGraph[r][startCol:c]))
				if err != nil {
					return 0, errors.New("unable to convert part number to int")
				}
				partNumSum += partNum
			}
			validPartNum = false
			startCol = -1
		}
		// captures edge case where last character(s) are digit(s)
		if startCol > -1 && validPartNum {
			partNum, err := strconv.Atoi(string(schemGraph[r][startCol:]))
			if err != nil {
				return 0, errors.New("unable to convert part number to int")
			}
			partNumSum += partNum
		}
	}
	return partNumSum, nil
}
