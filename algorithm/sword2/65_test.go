package sword2

import (
	"testing"
)

func minimumLengthEncoding(words []string) int {
	trie := &Trie65{}
	m := make(map[string]struct{})
	for _, w := range words {
		trie.Insert(w, m)
	}
	var ans int
	for k := range m {
		ans++
		ans += len(k)
	}
	return ans
}

type Trie65 struct {
	children [26]*Trie65
	isEnd    bool
	prefix   int
}

func (t *Trie65) Insert(word string, m map[string]struct{}) {
	var node = t
	for i := len(word) - 1; i >= 0; i-- {
		if node.children[word[i]-'a'] == nil {
			node.children[word[i]-'a'] = &Trie65{}
		}
		node = node.children[word[i]-'a']
		node.prefix++
		if node.isEnd {
			node.isEnd = false
			delete(m, word[i:])
		}
	}
	for _, v := range node.children {
		if v != nil {
			return
		}
	}
	node.isEnd = true
	m[word] = struct{}{}
}

/*
["time","time","time","time"]
*/

func Test65(t *testing.T) {
	var words = []string{"time", "time", "time", "time"}
	t.Log(minimumLengthEncoding(words))
}
