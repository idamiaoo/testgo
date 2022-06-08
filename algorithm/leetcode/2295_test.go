package leetcode

func arrayChange(nums []int, operations [][]int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		m[v] = i
	}
	for _, operation := range operations {
		index := m[operation[0]]
		delete(m, operation[0])
		m[operation[1]] = index
	}
	for k, v := range m {
		nums[v] = k
	}
	return nums
}
