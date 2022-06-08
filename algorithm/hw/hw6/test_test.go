package hw6

import (
	"strings"
	"testing"
)

func splitWord(s string) string {
	var words []string
	var isStart bool
	var word []byte
	for i := 0; i < len(s); i++ {
		if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') {
			if isStart {
				word = append(word, s[i])
			} else {
				if s[i] >= 'a' && s[i] <= 'z' {
					word = append(word, 'A'+(s[i]-'a'))
				} else {
					word = append(word, s[i])
				}
				isStart = true
			}
		} else {
			if isStart {
				if len(word) > 0 {
					words = append(words, string(word))
					word = nil
				}
				isStart = false
			}
		}
	}
	if isStart {
		words = append(words, string(word))
	}

	if len(words) == 0 {
		return "."
	}
	return strings.Join(words, " ") + "."
}

func TestSplitWord(t *testing.T) {
	var cases = []struct {
		src    string
		result string
	}{
		{
			src:    "who love?, Solo..",
			result: "Who Love Solo.",
		},
		{
			src:    "?????????",
			result: ".",
		},
		{
			src:    "?????????aa",
			result: "Aa.",
		},
		{
			src:    "   12Aa",
			result: "Aa.",
		},
		{
			src:    "   12Aa",
			result: "Aa.",
		},
	}

	for _, v := range cases {
		if splitWord(v.src) != v.result {
			t.Fatal(v.src, splitWord(v.src), v.result)
		}
	}
}
