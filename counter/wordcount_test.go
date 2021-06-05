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

func TestShouldCountUTF8(t *testing.T) {
	assert.Equal(t, WordCount("Japan 日本"),
		map[string]int{
			"Japan": 1, "日本": 1})
}

func TestShouldCountByteCode(t *testing.T) {
	assert.Equal(t, WordCount("\xe6\x97\xa5"),
		map[string]int{"日": 1})
}

func TestShouldCountRawStringLiteral(t *testing.T) {
	assert.Equal(t, WordCount("`\xe6`"),
		map[string]int{"`\xe6`": 1})
}

func TestShouldCountNonWhitespaceC0Controls(t *testing.T) {
	assert.Equal(t, WordCount(" a"),
		map[string]int{"\x01\x02\x03\x04\x05\x06\a\b\x0e\x0f\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\x1a\x1b\x1c\x1d\x1e\x1f\u007f": 1,
			"a": 1})
}

func TestShouldCountNonWhitespaceC1Controls(t *testing.T) {
	assert.Equal(t, WordCount(" a"),
		map[string]int{"\u0080\u0081\u0082\u0083\u0084\u0086\u0087\u0088\u0089\u008a\u008b\u008c\u008d\u008e\u008f\u0090\u0091\u0092\u0093\u0094\u0095\u0096\u0097\u0098\u0099\u009a\u009b\u009c\u009d\u009e\u009f": 1,
			"a": 1})
}

func TestShouldHandleWhitespaceZXControls(t *testing.T) {
	assert.Equal(t, WordCount("	              ​    　 a"),
		map[string]int{"\u200b": 1,
			"a": 1})
}

func TestShouldHandleByteOrderMarks(t *testing.T) {
	assert.Equal(t, WordCount("￾ a"),
		map[string]int{"\ufffe": 1,
			"a": 1})
}

func TestShouldHandleASuperStringRecommendedByVMware(t *testing.T) {
	assert.Equal(t, WordCount("表ポあA鷗Œé Ｂ逍Üßªąñ丂㐀𠀀 a"),
		map[string]int{"表ポあA鷗Œé": 1, "Ｂ逍Üßªąñ丂㐀𠀀": 1,
			"a": 1})
}
