package leetcode

import "testing"

func combine(n int, k int) [][]int {
	var (
		ans          [][]int
		path         []int
		backtracking func(index int)
	)
	backtracking = func(index int) {
		if len(path) == k {
			ans = append(ans, append([]int{}, path...))
			return
		}
		for i := index; i <= n; i++ {
			path = append(path, i)
			backtracking(i + 1)
			path = path[:len(path)-1]
		}
	}
	backtracking(1)
	return ans
}

func permutation(nums []int) [][]int {
	var (
		ans          [][]int
		path         []int
		backtracking func()
		choice       = make(map[int]int)
	)
	backtracking = func() {
		if len(path) == len(nums) {
			ans = append(ans, append([]int{}, path...))
			return
		}
		for i := 0; i < len(nums); i++ {
			if choice[i] == 1 {
				continue
			}

			choice[i] = 1
			path = append(path, nums[i])
			backtracking()
			path = path[:len(path)-1]
			choice[i] = 0
		}

	}
	backtracking()
	return ans
}

func quick77(nums []int, left, right int) []int {
	if left < right {
		p := partition(nums, left, right)
		quick77(nums, left, p-1)
		quick77(nums, p+1, right)
	}
	return nums
}

func partition(nums []int, left, right int) int {
	p := nums[right]
	i := left - 1
	for j := left; j < right; j++ {
		if nums[j] <= p {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[right] = nums[right], nums[i+1]
	return i + 1
}

func TestPermutation(t *testing.T) {
	t.Log(permutation([]int{1, 2, 3}))
}

func TestCombine(t *testing.T) {
	t.Log(combine(4, 2))
	t.Log(quick77([]int{3, 9, 6, 8, 4, 7, 9, 3, 4, 7, 5}, 0, 10))
}
