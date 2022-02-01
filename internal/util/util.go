package util

import (
	"fmt"
	"log"
	"sort"
	"time"
)

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
func SortMapByValue(m map[string]int) PairList {
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

func Track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func Duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}
