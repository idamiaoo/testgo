package leetcode

func canJump(nums []int) bool {
	var maxDistance int
	for i, v := range nums {
		if maxDistance < i {
			return false
		}
		if maxDistance < i+v {
			maxDistance = i + v
		}
		if maxDistance >= len(nums)-1 {
			return true
		}
	}
	return false
}
