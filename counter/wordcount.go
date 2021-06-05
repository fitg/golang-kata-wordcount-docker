package wordcount

import (
	"strings"
)

// WordCount is a function counting number of words in a sentence
// example WordCount("foo bar") would return
// foo: 1, bar: 1
func WordCount(sentence string) map[string]int {
	wordCount := make(map[string]int)

	words := strings.Fields(sentence)
	for _, word := range words {
		wordCount[word]++
	}

	return wordCount
}
