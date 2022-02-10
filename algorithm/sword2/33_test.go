package sword2

import (
	"testing"
)

func groupAnagrams(strs []string) [][]string {
	count := make([][26]int, len(strs))
	for i, str := range strs {
		cnt := [26]int{}
		for _, v := range []byte(str) {
			cnt[v-'a']++
		}
		count[i] = cnt
	}
	temp := make(map[[26]int][]int)
	for i, v := range count {
		if _, ok := temp[v]; ok {
			temp[v] = append(temp[v], i)
		} else {
			temp[v] = []int{i}
		}
	}
	var res [][]string
	for _, v := range temp {
		var sub []string
		for _, i := range v {
			sub = append(sub, strs[i])
		}
		res = append(res, sub)
	}
	return res
}

func Test33(t *testing.T) {
	strs := []string{"a"}
	t.Log(groupAnagrams(strs))
}
