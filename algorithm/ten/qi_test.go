package ten

import "testing"

func TestRev(t *testing.T) {
	n1 := &Node{
		Val: 1,
	}
	n2 := &Node{
		Val: 2,
	}
	n3 := &Node{
		Val: 3,
	}
	n4 := &Node{
		Val: 4,
	}
	n5 := &Node{
		Val: 5,
	}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5

	n := revList(n1)
	t.Log(n)
}

func TestMerge(t *testing.T) {
	n1 := &Node{
		Val: 1,
	}
	n2 := &Node{
		Val: 2,
	}
	n3 := &Node{
		Val: 3,
	}
	n4 := &Node{
		Val: 4,
	}
	n5 := &Node{
		Val: 5,
	}
	n6 := &Node{
		Val: 6,
	}
	n7 := &Node{
		Val: 7,
	}

	n1.Next = n3
	n3.Next = n5
	n5.Next = n7
	n2.Next = n4
	n4.Next = n6

	n := mergeList(n1, n2)

	t.Log(n)
}

func TestUp(t *testing.T) {

}
