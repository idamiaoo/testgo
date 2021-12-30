package leetcode

import (
	"fmt"
	"testing"
)

func numFriendRequests(ages []int) int {
	var res int
	if len(ages) == 1 {
		return res
	}
	sort825(ages, 0, len(ages)-1)
	fmt.Println(ages)
	for x := len(ages) - 1; x > 0; x-- {
		i := firstLarge(ages, 0, x-1, ages[x]/2+8)
		if i >= 0 {
			fmt.Println(">", ages[x], i, ages[i], float64(ages[x])*0.5+7.0, x-i)
			res += x - i
		}
		i = equal(ages, x-1, ages[x])
		if i >= 0 {
			fmt.Println("=", ages[x], x-i)
			res += x - i
		}

	}
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
	for j := left; j < right; j++ {
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

func equal(nums []int, i, base int) int {
	res := -1
	for i >= 0 && nums[i] == base {
		res = i
		i--
	}
	return res
}

func Test825(t *testing.T) {
	ages := []int{73, 106, 39, 6, 26, 15, 30, 100, 71, 35, 46, 112, 6, 60, 110}
	t.Log(numFriendRequests(ages))
}
