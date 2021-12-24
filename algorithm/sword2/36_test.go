package sword2

import (
	"strconv"
	"testing"
)

func evalRPN(tokens []string) int {
	s := newStack()
	var v, a, b int
	for _, token := range tokens {
		switch token {
		case "+":
			s.Put(s.Pop() + s.Pop())
		case "-":
			s.Put(-s.Pop() + s.Pop())
		case "*":
			s.Put(s.Pop() * s.Pop())
		case "/":
			a, b = s.Pop(), s.Pop()
			s.Put(b / a)
		default:
			v, _ = strconv.Atoi(token)
			s.Put(v)
		}
	}
	return s.Pop()
}

//["2","1","+","3","*"]

func Test36(t *testing.T) {
	tokens := []string{"2", "1", "+", "3", "*"}
	t.Log(evalRPN(tokens))
}
