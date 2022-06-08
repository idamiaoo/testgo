package leetcode

import "fmt"

type TextEditor struct {
	head   *editorNode
	tail   *editorNode
	cursor *editorNode
}

type editorNode struct {
	v    rune
	prev *editorNode
	next *editorNode
}

func Constructor2296() TextEditor {
	cursor := &editorNode{
		v: '|',
	}
	head := &editorNode{
		v:    0,
		next: cursor,
	}
	tail := &editorNode{
		v:    0,
		prev: cursor,
	}
	cursor.prev = head
	cursor.next = tail
	return TextEditor{
		head:   head,
		tail:   tail,
		cursor: cursor,
	}
}

func (this *TextEditor) AddText(text string) {
	var head, tail *editorNode
	for _, v := range text {
		node := &editorNode{
			v: v,
		}
		if head == nil {
			head = node
			tail = node
		} else {
			tail.next = node
			node.prev = tail
			tail = node
		}
	}
	this.cursor.prev.next = head
	head.prev = this.cursor.prev
	this.cursor.prev = tail
	tail.next = this.cursor
	fmt.Println(this.String())
}

func (this *TextEditor) DeleteText(k int) int {
	temp := this.cursor.prev
	v := k
	for k > 0 && temp != this.head {
		temp = temp.prev
		k--
	}
	temp.next = this.cursor
	this.cursor.prev = temp
	fmt.Println(this.String())
	return v - k
}

func (this *TextEditor) CursorLeft(k int) string {
	temp := this.cursor.prev
	for k > 0 && temp != this.head {
		temp = temp.prev
		k--
	}
	this.cursor.prev.next = this.cursor.next
	this.cursor.next.prev = this.cursor.prev
	temp.next = this.cursor
	this.cursor.prev = temp
	temp = this.cursor.prev
	i := 0
	s := ""
	for i < 10 && temp != this.head {
		s = string(temp.v) + s
		temp = temp.prev
		i++
	}
	fmt.Println(this.String())
	return s
}

func (this *TextEditor) CursorRight(k int) string {
	fmt.Println("right", k, this.String())

	temp := this.cursor.next
	for k > 0 && temp != this.tail {
		temp = temp.next
		k--
	}
	this.cursor.prev.next = this.cursor.next
	this.cursor.next.prev = this.cursor.prev

	this.cursor.prev = temp.prev
	temp.prev = this.cursor
	this.cursor.next = temp

	temp = this.cursor.prev
	i := 0
	s := ""
	for i < 10 && temp != this.head {
		s = string(temp.v) + s
		temp = temp.prev
		i++
	}
	fmt.Println("right", k, this.String())
	return s
}

func (this *TextEditor) String() string {
	s := ""
	temp := this.head.next
	for temp != this.tail {
		s += string(temp.v)
		temp = temp.next
	}
	return s
}
