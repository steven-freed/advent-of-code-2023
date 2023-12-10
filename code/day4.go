package code

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

func getCardMatches(row string) map[string]int {
	checkingPicks := false
	matches := map[string]int{}
	for j := 0; j < len(row); j++ {
		k := j
		if !isDigit(rune(row[j])) {
			if row[j] == 124 { // 124 == '|'
				checkingPicks = true
			}
			continue
		}
		for k < len(row) && isDigit(rune(row[k])) {
			k++
		}
		if !checkingPicks {
			matches[row[j:k]] = 0
		} else if _, isWinner := matches[row[j:k]]; isWinner {
			matches[row[j:k]] += 1
		}
		j = k
	}
	return matches
}

func Scratchcards_Part2(input string) (int, error) {
	lines := strings.Split(input, "\n")
	cardCopies := make([]int, len(lines))
	for i := range cardCopies {
		cardCopies[i] = 1
	}
	for i := 0; i < len(lines); i++ {
		numLineStr := strings.TrimSpace(lines[i][strings.Index(lines[i], ":")+1:])
		cardNumStr := strings.TrimSpace(lines[i][strings.Index(lines[i], "Card")+len("Card") : strings.Index(lines[i], ":")])
		cardNum, err := strconv.Atoi(cardNumStr)
		if err != nil {
			return 0, errors.New("unable to parse card number")
		}
		rowMatchCount := 0
		for _, cnt := range getCardMatches(numLineStr) {
			rowMatchCount += cnt
		}
		for j := 1; j <= rowMatchCount; j++ {
			cardCopies[cardNum-1+j] += cardCopies[cardNum-1]
		}
	}
	tot := 0
	for _, n := range cardCopies {
		tot += n
	}
	return tot, nil
}

func Scratchcards_Part1(input string) (int, error) {
	winnings := 0
	lines := strings.Split(input, "\n")
	for i := 0; i < len(lines); i++ {
		numLineStr := strings.TrimSpace(lines[i][strings.Index(lines[i], ":")+1:])
		cardWinnings := 0
		for _, pick := range getCardMatches(numLineStr) {
			cardWinnings += pick
		}
		winnings += int(math.Pow(2, float64(cardWinnings)-1))
	}
	return winnings, nil
}
