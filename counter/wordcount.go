package wordcount

import (
	"strings"
)

// WordCount is a function counting number of words in a sentence
// it uses ASCII whitespace (0x20) to separate words
// https://en.wikipedia.org/wiki/Whitespace_character
// example WordCount("foo bar") would return
// map[bar:1 foo:1]
func WordCount(sentence string) map[string]int {
	wordCount := make(map[string]int)

	words := strings.Fields(sentence)
	for _, word := range words {
		wordCount[word]++
	}

	return wordCount
}
