package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type WordCount struct {
	Word   string
	Count  int
}
type WordList []WordCount

func (p WordList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p WordList) Len() int           { return len(p) }
func (p WordList) Less(i, j int) bool { return p[i].Count < p[j].Count }

func sortByWordCount(wordFreq map[string]int) WordList {
	//list := make(WordList, len(wordFreq))
	list := make(WordList, 0)
	for word, count := range wordFreq {
                entry := WordCount{word, count}
		list = append(list, entry)
	}
	sort.Sort(sort.Reverse(list))
	return list
}


func countWords(text string) WordList {
	// Run the regular expression on string to replace everything but
	// alpha-numericals with space
	var regExp = regexp.MustCompile("[[:^alpha:]]")
	cleanedTxt := regExp.ReplaceAllString(text, " ")

	// Build the Map(string, int) for this string
	cleanedTxt = strings.ToLower(cleanedTxt)
	wordArr := strings.Fields(cleanedTxt)
	wordFreq := make(map[string]int)
	for _, v := range wordArr {
		wordFreq[v] = wordFreq[v] + 1
	}

	//Sort the map on Value
	sortedByCount := sortByWordCount(wordFreq)
        return sortedByCount
}

func main() {

	if (len(os.Args)) != 2 {
		fmt.Println("Please provide one filename as argument.")
		return
	}

	// How can I typecast dat into String without using a new var
	dat, err := ioutil.ReadFile(os.Args[1])
	check(err)
	txtData := string(dat)

        sortedByCount := countWords(txtData)

	for _, pl := range sortedByCount {
		fmt.Println(pl.Count, pl.Word)
	}
}
