package code

import (
	"fmt"
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

func GearRatios_Part1(input string) (int, error) {
	// convert input to graph
	rows := strings.Split(input, "\n")
	rowCount := len(rows)
	var colCount int
	if len(rows) > 0 {
		colCount = len(rows[0])
	}
	schemGraph := make([][]rune, rowCount)
	markedGraph := make([][]rune, rowCount)
	for i, row := range rows {
		schemGraphRow := make([]rune, colCount)
		markedGraphRow := make([]rune, colCount)
		for j, col := range row {
			schemGraphRow[j] = col
			markedGraphRow[j] = col
		}
		schemGraph[i] = schemGraphRow
		markedGraph[i] = markedGraphRow
	}
	// create marked rows graph
	// each partnumber is transformed so that the first (left most) digit is the len of the number
	// and each subsequent digit from left to right is the offset from the first (left most) digit
	for i, _ := range schemGraph {
		partNumLen := 0
		startCol := -1
		for j, _ := range schemGraph[i] {
			if isDigit(schemGraph[i][j]) {
				if startCol < 0 {
					startCol = j
				}
				markedGraph[i][j] = rune(j - startCol + 48)
				partNumLen += 1
				continue
			}
			if startCol > -1 {
				markedGraph[i][startCol] = rune(partNumLen + 48)
				markedGraph[i][j-1] = rune(j - 1 - startCol + 48)
			}
			partNumLen = 0
			startCol = -1
		}
	}
	prettyPrint(markedGraph)
	// sum part numbers
	// for i, _ := range schemGraph {
	// 	for j, _ := range schemGraph[i] {
	// 		if schemGraph[i][j] >= 48 && schemGraph[i][j] <= 57 {
	// 			if (i-1 > -1 && schemGraph[i-1][j] < 48 && schemGraph[i-1][j] > 57 && schemGraph[i-1][j] != '.') ||
	// 				(i+1 < rowCount && schemGraph[i+1][j] < 48 && schemGraph[i+1][j] > 57 && schemGraph[i+1][j] != '.') ||
	// 				(j-1 > -1 && schemGraph[i][j-1] < 48 && schemGraph[i][j-1] > 57 && schemGraph[i][j-1] != '.') ||
	// 				(j+1 < colCount && schemGraph[i][j+1] < 48 && schemGraph[i][j+1] > 57 && schemGraph[i][j+1] != '.') {

	// 			}
	// 		}
	// 	}
	// }
	return 0, nil
}
