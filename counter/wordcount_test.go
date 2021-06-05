package wordcount

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldReturnNothingForEmptyString(t *testing.T) {
	assert.Equal(t, WordCount(""), map[string]int{})
}

func TestShouldReturnNothingForWhitespace(t *testing.T) {
	assert.Equal(t, WordCount("        "), map[string]int{})
}

func TestShouldCountOneWord(t *testing.T) {
	assert.Equal(t, WordCount("foo"), map[string]int{"foo": 1})
}

func TestShouldCountOneWordMultipleTimes(t *testing.T) {
	assert.Equal(t, WordCount("foo foo"), map[string]int{"foo": 2})
}

func TestShouldCountTwoWords(t *testing.T) {
	assert.Equal(t, WordCount("foo bar"), map[string]int{"foo": 1, "bar": 1})
}

func TestShouldCountMultipleWordsMultipleTimes(t *testing.T) {
	assert.Equal(t, WordCount("foo bar yo foo yo test yarrr"),
		map[string]int{
			"foo": 2, "bar": 1, "yo": 2,
			"test": 1, "yarrr": 1})
}
