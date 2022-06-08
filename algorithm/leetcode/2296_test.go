package leetcode

type TextEditor struct {
	head *editorNode
	tail *editorNode
}

type editorNode struct {
	v    rune
	prev *editorNode
	next *editorNode
}

func Constructor2296() TextEditor {
	node := &editorNode{
		v: '|',
	}
	head := &editorNode{
		v:    0,
		next: node,
	}
	tail := &editorNode{
		v:    0,
		prev: node,
	}
	node.prev = head
	node.next = tail
	return TextEditor{
		head: head,
		tail: tail,
	}
}

func getNode(text string) {
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
	for this.head.next != nil {

	}
}

func (this *TextEditor) DeleteText(k int) int {

}

func (this *TextEditor) CursorLeft(k int) string {

}

func (this *TextEditor) CursorRight(k int) string {

}
