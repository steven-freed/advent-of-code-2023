package code

const (
	GAME  = "Game"
	COLON = ":"
)

func removeTrailingNewLine(input string) string {
	return input[:len(input)-1]
}

func isDigit(input rune) bool {
	return input >= 48 && input <= 57
}

func find(haystack string, needle string) int {
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}

func rfind(haystack string, needle string) int {
	for i := len(haystack) - len(needle); i >= 0; i-- {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}
