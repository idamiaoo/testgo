package sword2

import (
	"testing"
)

type Trie struct {
	count  int
	prefix int
	nodes  [26]*Trie
}

/** Initialize your data structure here. */
func Constructor62() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	if word == "" {
		this.count = 1
		return
	}
	node := this.nodes[word[0]-'a']
	if node == nil {
		node = &Trie{}
		this.nodes[word[0]-'a'] = node
	}
	node.prefix++
	node.Insert(word[1:])
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	if word == "" {
		return this.count == 1
	}
	node := this.nodes[word[0]-'a']
	if node == nil {
		return false
	}
	return node.Search(word[1:])
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	if prefix == "" {
		for _, v := range this.nodes {
			if v != nil {
				return true
			}
		}
		return false
	}
	node := this.nodes[prefix[0]-'a']
	if node == nil {
		return false
	}
	return node.StartsWith(prefix[1:])
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);

["Trie","insert","search","search","insert","search","insert","search"]
[[],["abc"],["abc"],["ab"],["ab"],["ab"],["ab"],["ab"]]
*/

func Test62(t *testing.T) {
	trie := Constructor62()

	trie.Insert("apple")
	t.Log(trie.Search("apple"))
	t.Log(trie.Search("app"))
}
