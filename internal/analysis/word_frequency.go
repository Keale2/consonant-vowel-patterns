package analysis

import (
	"fmt"
	"strings"
	"unicode"
)

type AnalysisResults struct {
	word string
	g, y float32
}

func getLetterType(letter rune) rune {
	letter = unicode.ToLower(letter)
	switch letter {
	case 'a', 'e', 'i', 'o', 'u':
		return 'O'
	case 'b', 'c', 'd', 'f', 'g', 'h', 'j', 'k', 'l', 'm', 'n', 'p', 'q', 'r', 's', 't', 'v', 'w', 'x', 'z':
		return 'L'
	case 'y':
		return 'Y'
	default:
		return letter
	}
}

func Analyze(words []string) map[string]int {
	patternMap := map[string]int{}

	fmt.Println("DOIN AN ANALYSIS")
	// Loop through all 5-letter words
	for i := range words {
		word := words[i]
		pattern := strings.Map(getLetterType, word)

		patternMap[pattern]++
	}
	//fmt.Println(patternMap)
	return patternMap
}
