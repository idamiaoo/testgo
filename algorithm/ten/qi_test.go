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

func permutation(str string) []string {
	var res []string
	var backtrack func([]byte, int, []byte)
	backtrack = func(src []byte, i int, s []byte) {
		if i < 0 || i >= len(src) || src[i] == ' ' {
			return
		}
		s = append(s, src[i])
		if len(s) == len(src) {
			res = append(res, string(s))
			return
		}
		src[i] = ' '
		for i := 0; i < len(src); i++ {
			backtrack(src, i, s)
		}
		src[i] = s[len(s)-1]
		s = s[:len(s)-1]
	}
	for i := range str {
		backtrack([]byte(str), i, nil)
	}
	return res
}

func TestPermutation(t *testing.T) {
	t.Log(permutation("abc"))
}
