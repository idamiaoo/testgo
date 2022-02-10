package leetcode

import (
	"sort"
)

func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false
	}
	sort.Ints(hand)
	cnt := make(map[int]int)
	for _, v := range hand {
		cnt[v]++
	}
	for _, v := range hand {
		if cnt[v] == 0 {
			continue
		}
		for i := v; i < v+groupSize; i++ {
			if cnt[i] == 0 {
				return false
			}
			cnt[i]--
		}
	}
	return true
}
