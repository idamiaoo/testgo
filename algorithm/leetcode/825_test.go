package leetcode

func numFriendRequests(ages []int) int {
	var res int
	if len(ages) == 1 {
		return res
	}
	sort825(ages, 0, len(ages)-1)
	return res
}

func sort825(nums []int, left, right int) {
	if left < right {
		p := partition825(nums, left, right)
		sort825(nums, left, p-1)
		sort825(nums, p+1, right)
	}
}

func partition825(nums []int, left, right int) int {
	b := nums[right]
	i := left - 1
	for j := 0; j < right; j++ {
		if nums[j] < b {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	i++
	nums[i], nums[right] = nums[right], nums[i]
	return i
}

func firstLarge(nums []int, left, right, base int) int {
	if base > nums[right] {
		return -1
	}
	if base < nums[left] {
		return left
	}
	if left == right {
		return right
	}
	v := (left + right) / 2
	if nums[v] >= base {
		return firstLarge(nums, left, v, base)
	} else {
		return firstLarge(nums, v+1, right, base)
	}
}
