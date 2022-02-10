package sword2_test

import (
	"testing"
)

type Node struct {
	Key  int
	Val  int
	Prev *Node
	Next *Node
}

type LRUCache struct {
	Capacity int
	M        map[int]*Node
	Head     *Node
	Tail     *Node
}

func Constructor1(capacity int) LRUCache {
	head, tail := &Node{}, &Node{}
	head.Next = tail
	tail.Prev = head
	return LRUCache{
		Capacity: capacity,
		M:        make(map[int]*Node),
		Head:     head,
		Tail:     tail,
	}
}

func (this *LRUCache) RemoveNode(node *Node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (this *LRUCache) AddNodeToHead(node *Node) {
	node.Next = this.Head.Next
	node.Prev = this.Head
	node.Next.Prev = node
	this.Head.Next = node
}

func (this *LRUCache) MoveNodeToHead(node *Node) {
	this.RemoveNode(node)
	this.AddNodeToHead(node)
}

func (this *LRUCache) Get(key int) int {
	var val = -1
	if node, ok := this.M[key]; ok {
		val = node.Val
		this.MoveNodeToHead(node)
	}
	return val
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.M[key]; ok {
		node.Val = value
		this.MoveNodeToHead(node)
	} else {
		if len(this.M) >= this.Capacity {
			delete(this.M, this.Tail.Prev.Key)
			this.RemoveNode(this.Tail.Prev)
		}
		node := &Node{Key: key, Val: value}
		this.M[key] = node
		this.AddNodeToHead(node)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func Test31(t *testing.T) {
	lru := Constructor1(2)

	lru.Put(1, 1)
	lru.Put(2, 2)
	lru.Put(3, 3)

	t.Log(lru.Get(3))
}
