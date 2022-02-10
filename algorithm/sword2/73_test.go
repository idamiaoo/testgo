package sword2

func minEatingSpeed(piles []int, h int) int {
	var min, max, mid int
	for _, v := range piles {
		if v > max {
			max = v
		}
	}
	min = 1
	for min <= max {
		mid = min + ((max - min) >> 1)
		if countTime(piles, mid) <= h {
			if mid == 1 || countTime(piles, mid-1) > h {
				return mid
			}
			max = mid - 1
		} else {
			min = mid + 1
		}
	}
	return -1
}

func countTime(piles []int, k int) int {
	var t int
	for _, v := range piles {
		t += (v + k - 1) / k
	}
	return t
}
