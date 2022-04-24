package leetcode

func hammingDistance(x int, y int) int {
	var ans int
	for x > 0 && y > 0 {
		if x&1 != y&1 {
			ans++
		}
		x >>= 1
		y >>= 1
	}
	for x > 0 {
		if x&1 == 1 {
			ans++
		}
		x >>= 1
	}

	for y > 0 {
		if y&1 == 1 {
			ans++
		}
		y >>= 1
	}
	return ans
}
