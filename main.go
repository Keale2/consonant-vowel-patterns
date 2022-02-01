package main

import (
	"github.com/Keale2/consonant-vowel-patterns/cmd"
)

func main() {
	cmd.Execute()
}

// https://groups.google.com/g/golang-nuts/c/FT7cjmcL7gw
// A data structure to hold a key/value pair
// It looks like this was missing i++

// func main() {
// 	defer duration(track("main"))
// 	if len(os.Args) <= 1 {
// 		fmt.Printf("USAGE : %s <target_filename> \n", os.Args[0])
// 		os.Exit(0)
// 	}

// 	// Write all 5-letter words to a new file
// 	printLines("./five.txt", fiveLetterWords)

// 	// Write patterns to file sorted by frequency
// 	printLines("./pattern-frequency.txt", sortMapByValue(patternMap).String())
// }
