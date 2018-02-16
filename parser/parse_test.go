package main

import (
	"testing"
	"github.com/bmizerany/assert"
	"strings"
)

func TestParseWords(t *testing.T) {
	cases := []map[string][]string{
		{
			"input": {"My name is mohan", "I write code"},
			"expected": {"My","name","is","mohan","I","write","code"},
		},
	}

	for _, c := range cases {
		out := ParseWords(c["input"])
		assert.Equal(t, c["expected"], out)
	}
}

func TestParseLines(t *testing.T) {
	inputStr := `My name is Mohan
I write code`

	expected := []string{"My name is Mohan\n", "I write code"}
	out := ParseLines(strings.NewReader(inputStr))
	assert.Equal(t, expected, out)
}
