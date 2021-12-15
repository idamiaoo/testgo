package sword2

import (
	"strings"
	"testing"
)

func maxProduct(words []string) int {
	var res int
	temp := make([]int, len(words))
	for i, word := range words {
		for _, w := range []byte(word) {
			temp[i] |= 1 << (w - 'a')
		}
	}

	for i := 0; i < len(temp)-1; i++ {
		for j := i + 1; j < len(temp); j++ {
			if temp[i]&temp[j] == 0 {
				if len(words[i])*len(words[j]) > res {
					res = len(words[i]) * len(words[j])
				}
			}
		}
	}
	return res
}

// 108
func Test5(t *testing.T) {
	var words = []string{"abcw", "baz", "foo", "bar", "fxyz", "abcdef"}
	t.Log(maxProduct(words))
}

func TestSplit(t *testing.T) {
	v := strings.SplitN("create_time.gt:2021-12-11 07:47:42", ":", 2)
	t.Log(len(v), v)
	t.Log(strings.Contains("pmnnsfpjhcpwkmcdbfvlxngcdmejcvxpktilhywpitxwsfavpmlmdczpyilnhakjkoscgyxllac", "t"))
}
