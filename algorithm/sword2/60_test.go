package sword2

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	counts := make([]counter, 0, len(nums))
	for c, v := range m {
		counts = append(counts, counter{
			v: v,
			c: c,
		})
	}
	topK(counts, k)
	var ans []int
	for i := 0; i < k; i++ {
		ans = append(ans, counts[i].c)
	}
	return ans
}

type counter struct {
	v int
	c int
}

func topK(nums []counter, k int) {
	if len(nums) == k {
		return
	}
	p := partition60(nums, 0, len(nums)-1)
	if p+1 == k {
		return
	}
	if p+1 > k {
		topK(nums[:p], k)
	} else {
		topK(nums[p+1:], k-p-2)
	}
}

func partition60(nums []counter, left, right int) int {
	b := nums[right].v
	i := left - 1
	for j := 0; j < right; j++ {
		if nums[j].v > b {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	i++
	nums[i], nums[right] = nums[right], nums[i]
	return i
}
