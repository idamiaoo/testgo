package sword2

import (
	"strings"
	"testing"
)

func replaceWords(dictionary []string, sentence string) string {
	t := &trie{}
	for _, v := range dictionary {
		t.insert(v)
	}
	var ans []string
	for _, v := range strings.Split(sentence, " ") {
		ans = append(ans, t.replace(v))
	}
	return strings.Join(ans, " ")
}

type trie struct {
	children [26]*trie
	isEnd    bool
}

func (t *trie) insert(word string) {
	var node = t
	for _, w := range word {
		if node.children[w-'a'] == nil {
			node.children[w-'a'] = &trie{}
		}
		node = node.children[w-'a']
	}
	node.isEnd = true
}

func (t *trie) replace(word string) string {
	var node = t
	var ans []rune
	for _, w := range word {
		node = node.children[w-'a']
		if node == nil {
			break
		}
		ans = append(ans, w)
		if node.isEnd {
			return string(ans)
		}
	}
	return word
}

/*
例 1：

输入：dictionary = ["cat","bat","rat"], sentence = "the cattle was rattled by the battery"
输出："the cat was rat by the bat"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/UhWRSj
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func Test63(t *testing.T) {
	var dictionary = []string{"cat", "bat", "rat"}
	var sentence = "the cattle was rattled by the battery"
	t.Log(replaceWords(dictionary, sentence))
}
