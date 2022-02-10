package sword2_test

import (
	"testing"
)

type MagicDictionary struct {
	trie *Trie
}

/** Initialize your data structure here. */
func Constructor64() MagicDictionary {
	return MagicDictionary{trie: &Trie{}}
}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	for _, w := range dictionary {
		this.trie.Insert(w)
	}
}

func (this *MagicDictionary) Search(searchWord string) bool {
	return this.MagicSearch(this.trie, searchWord, 1)
}
func (this *MagicDictionary) MagicSearch(trie *Trie, searchWord string, change int) bool {
	if trie == nil {
		return false
	}
	if trie.isEnd && searchWord == "" && change == 0 {
		return true
	}

	if len(searchWord) > 0 {
		var found bool
		for ii, n := range trie.children {
			if ii == int(searchWord[0]-'a') {
				found = this.MagicSearch(n, searchWord[1:], change)
			} else {
				if change > 0 {
					found = this.MagicSearch(n, searchWord[1:], change-1)
				}
			}
			if found {
				break
			}
		}
		return found
	}
	return false
}

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func (t *Trie) Insert(word string) {
	var node = t
	for _, v := range word {
		v -= 'a'
		if node.children[v] == nil {
			node.children[v] = &Trie{}
		}
		node = node.children[v]
	}
	node.isEnd = true
}

func (t *Trie) SearchPrefix(prefix string) *Trie {
	var node = t
	for _, v := range prefix {
		v -= 'a'
		if node.children[v] == nil {
			return nil
		}
		node = node.children[v]
	}
	return node
}

func (t *Trie) Search(word string) bool {
	node := t.SearchPrefix(word)
	return node != nil && node.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	node := t.SearchPrefix(prefix)
	return node != nil
}

/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dictionary);
 * param_2 := obj.Search(searchWord);

["MagicDictionary", "buildDict", "search", "search", "search", "search"]
[[], [["hello","hallo","leetcode"]], ["hello"], ["hhllo"], ["hell"], ["leetcoded"]]

["MagicDictionary", "buildDict", "search", "search", "search", "search", "search", "search", "search", "search"]
[[], [["a","b"]], ["a"], ["b"], ["c"], ["d"], ["e"], ["f"], ["ab"], ["ba"]]

[null,null,true,true,true,true,true,true,false,false]
*/

func Test64(t *testing.T) {
	var dict = []string{"hello", "hallo", "leetcode"}
	m := Constructor64()
	m.BuildDict(dict)
	t.Logf("%v,%v,%v,%v", m.Search("hello"), m.Search("hhllo"), m.Search("hell"), m.Search("leetcoded"))
	var a = "abc"
	t.Log(a[3:])
}
