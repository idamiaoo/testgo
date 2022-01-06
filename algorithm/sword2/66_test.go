package sword2

import (
	"testing"
)

type MapSum struct {
	sum      int
	isEnd    bool
	children [26]*MapSum
}

/** Initialize your data structure here. */
func Constructor66() MapSum {
	return MapSum{}
}

func (this *MapSum) Insert(key string, val int) {
	node := this
	for _, v := range key {
		v -= 'a'
		if node.children[v] == nil {
			node.children[v] = &MapSum{}
		}
		node = node.children[v]
	}
	node.isEnd = true
	node.sum = val
}

func (this *MapSum) Sum(prefix string) int {
	node := this
	for _, v := range prefix {
		v -= 'a'
		if node.children[v] == nil {
			return 0
		}
		node = node.children[v]
	}
	return sum(node)
}

func sum(root *MapSum) int {
	var ans int
	if root == nil {
		return ans
	}
	if root.isEnd {
		ans += root.sum
	}
	for _, v := range root.children {
		ans += sum(v)
	}
	return ans
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);

["MapSum", "insert", "sum", "insert", "insert", "sum"]
[[], ["apple",3], ["ap"], ["app",2], ["apple", 2], ["ap"]]

[null,null,3,null,null,7]
[null,null,3,null,null,4]
*/

func Test66(t *testing.T) {
	m := Constructor66()
	m.Insert("apple", 3)
	t.Log(m.Sum("ap"))
	m.Insert("app", 2)
	m.Insert("apple", 2)
	t.Log(m.Sum("ap"))

}
