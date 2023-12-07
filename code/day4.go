package code

import (
	"math"
	"strings"
)

func Scratchcards_Part1(input string) (int, error) {
	winnings := 0
	lines := strings.Split(input, "\n")
	for i := 0; i < len(lines); i++ {
		checkingPicks := false
		winningsCounts := map[string]int{}
		nums := strings.TrimSpace(lines[i][strings.Index(lines[i], ":"):])
		for j := 0; j < len(nums); j++ {
			k := j
			if !isDigit(rune(nums[j])) {
				if nums[j] == 124 { // 124 == '|'
					checkingPicks = true
				}
				continue
			}
			for k < len(nums) && isDigit(rune(nums[k])) {
				k++
			}
			if !checkingPicks {
				winningsCounts[nums[j:k]] = 0
			} else if _, isWinner := winningsCounts[nums[j:k]]; isWinner {
				winningsCounts[nums[j:k]] += 1
			}
			j = k
		}
		cardWinnings := 0
		for _, pick := range winningsCounts {
			cardWinnings += pick
		}
		winnings += int(math.Pow(2, float64(cardWinnings)-1))
	}
	return winnings, nil
}
