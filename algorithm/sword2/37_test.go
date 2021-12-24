package sword2

import (
	"testing"
)

func asteroidCollision(asteroids []int) []int {
	s := newStack()
	var i int
	for i = 0; i < len(asteroids); {
		if asteroids[i] > 0 || s.Size() == 0 || s.Top() < 0 {
			s.Put(asteroids[i])
			i++
		} else {
			n := s.Pop()
			if n+asteroids[i] > 0 {
				s.Put(n)
				i++
			} else if n+asteroids[i] == 0 {
				i++
			}
		}
	}
	return s.Range()
}

func Test37(t *testing.T) {
	var asteroids = []int{-2, -1, 1, 2}
	t.Log(asteroidCollision(asteroids))
}
