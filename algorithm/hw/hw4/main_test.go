package hw4

import "testing"

func TestXxx(t *testing.T) {
	var (
		gas  = []int{1, 2, 3, 4, 5}
		cost = []int{3, 4, 5, 1, 2}
	)
	t.Log(canCompleteCircuit(gas, cost))
}
