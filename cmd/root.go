package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"sort"
	"unicode"

	"github.com/SpawNKZ/kazdream/entities"
)

func App() {
	data, err := ioutil.ReadFile(entities.FileName)
	if err != nil {
		return
	}
	words := bytes.FieldsFunc(data, func(c rune) bool {
		return !unicode.IsLetter(c)
	})
	toLower(words)
	sortByteSlices(words)
	result := countWords(words)
	sortSlice(result)
	printWords(result)
}

func toLower(words [][]byte) {
	for i := range words {
		words[i] = bytes.ToLower(words[i])
	}
}

func sortByteSlices(words [][]byte) {
	sort.Slice(words, func(i, j int) bool {
		return bytes.Compare(words[i], words[j]) < 0
	})
}

func countWords(words [][]byte) []*entities.Word {
	var wordsWithAmount []*entities.Word
	var count = 1
	for i := range words {
		if i == len(words)-1 {
			break
		}
		if bytes.Compare(words[i], words[i+1]) == 0 {
			// if bytes.Equal(words[i], words[i+1]) {
			count++
		} else {
			w := &entities.Word{
				Value: words[i],
				Count: count,
			}
			wordsWithAmount = append(wordsWithAmount, w)
			count = 1
		}
	}
	return wordsWithAmount
}

func sortSlice(words []*entities.Word) {
	sort.Slice(words, func(i, j int) bool {
		return words[i].Count > words[j].Count
	})
}

func printWords(words []*entities.Word) {
	for i, word := range words {
		if i == entities.CountWordsForPrint {
			break
		}
		fmt.Printf("% 8d %s\n", word.Count, word.Value)
	}
}
