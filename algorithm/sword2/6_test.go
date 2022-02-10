package sword2

import (
	"testing"
)

func twoSum(numbers []int, target int) []int {
	for i, j := 0, len(numbers)-1; i < j; {
		if numbers[i]+numbers[j] > target {
			j--
		} else if numbers[i]+numbers[j] < target {
			i++
		} else {
			return []int{i, j}
		}
	}
	return nil
}

func Test6(t *testing.T) {
	numbers := []int{1, 2, 4, 6, 10}
	target := 8
	t.Log(twoSum(numbers, target))
}
