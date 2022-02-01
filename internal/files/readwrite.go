package files

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

func ReadFile(fileName string) ([]string, error) {
	fileBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
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

	return fiveLetterWords, nil
}

// Write each string in a slice to a new line in specified file
func WriteFile(filePath string, words []string) error {
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
