package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
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

// https://groups.google.com/g/golang-nuts/c/FT7cjmcL7gw
// A data structure to hold a key/value pair
// It looks like this was missing i++

type Pair struct {
	Key   string
	Value int
}

func (p Pair) String() string {
	str := fmt.Sprintf("%v %v", p.Key, p.Value)
	fmt.Println(str)
	return str
}

// A slice of Pairs that implements sort.Interface to sort by Value
type PairList []Pair

func (pl PairList) Swap(i, j int)      { pl[i], pl[j] = pl[j], pl[i] }
func (pl PairList) Len() int           { return len(pl) }
func (pl PairList) Less(i, j int) bool { return pl[i].Value < pl[j].Value }
func (pl PairList) String() []string {
	str := []string{}
	for _, p := range pl {
		str = append(str, fmt.Sprintf("%v %v", p.Key, p.Value))
	}

	return str
}

// A function to turn a map into a PairList, then sort and return it.
func sortMapByValue(m map[string]int) PairList {
	pl := make(PairList, len(m))
	i := 0
	for k, v := range m {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))

	// Could loop through p and create a string from each key/value pair...

	return pl
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

func main() {
	defer duration(track("main"))
	if len(os.Args) <= 1 {
		fmt.Printf("USAGE : %s <target_filename> \n", os.Args[0])
		os.Exit(0)
	}

	patternMap := map[string]int{}

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

	// Loop through all 5-letter words
	for i := range fiveLetterWords {
		word := fiveLetterWords[i]
		pattern := strings.Map(getLetterType, word)

		if _, ok := patternMap[pattern]; ok {
			patternMap[pattern]++
		} else {
			patternMap[pattern] = 1
		}

	}

	// Write all 5-letter words to a new file
	printLines("./five.txt", fiveLetterWords)

	// Write patterns to file sorted by frequency
	printLines("./pattern-frequency.txt", sortMapByValue(patternMap).String())

}
