package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}

// Write each string in a slice to a new line in specified file
func printLines(filePath string, words []string) error {
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	for _, word := range words {
		fmt.Fprintln(f, word) // print values to f, one per line
	}
	return nil
}

func main() {
	defer duration(track("main"))
	if len(os.Args) <= 1 {
		fmt.Printf("USAGE : %s <target_filename> \n", os.Args[0])
		os.Exit(0)
	}

	fileName := os.Args[1]

	fileBytes, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Read text file and make each line a slice
	sliceData := strings.Split(string(fileBytes), "\n")

	fiveLetterWords := []string{}

	// Look through all words and trim hidden characters
	for i := range sliceData {
		word := strings.TrimFunc(sliceData[i], func(r rune) bool {
			return !unicode.IsGraphic(r)
		})

		// Add all of the 5-letter words to a new slice
		if utf8.RuneCountInString(word) == 5 {
			fiveLetterWords = append(fiveLetterWords, word)
		}
	}

	// Write all 5-letter words to a new file
	printLines("./five.txt", fiveLetterWords)

}
