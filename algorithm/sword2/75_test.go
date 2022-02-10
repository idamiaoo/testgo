package sword2

func relativeSortArray(arr1 []int, arr2 []int) []int {
	order := make(map[int]int)
	for i, v := range arr2 {
		order[v] = i + 1
	}
	less := func(a, b int) bool {
		if order[a] == 0 && order[b] == 0 {
			return a < b
		}
		if order[a] != 0 && order[b] != 0 {
			return order[a] < order[b]
		}
		if order[a] == 0 {
			return false
		}
		return true
	}
	sort75(arr1, less, 0, len(arr1)-1)
	return arr1
}

func sort75(nums []int, less func(int, int) bool, left, right int) {
	if left < right {
		p := partition75(nums, less, left, right)
		sort75(nums, less, left, p-1)
		sort75(nums, less, p+1, right)
	}
}

func partition75(nums []int, less func(int, int) bool, left, right int) int {
	b := nums[right]
	i := left - 1
	for j := left; j < right; j++ {
		if less(nums[j], b) {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	i++
	nums[i], nums[right] = nums[right], nums[i]
	return i
}
