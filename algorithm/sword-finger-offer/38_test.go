package sword_finger_offer

import (
	"sort"
	"testing"
)

func permutation(s string) []string {
	res := make([]string, 0, 1)
	sorteds := []byte(s)
	sort.Slice(sorteds, func(i, j int) bool { return sorteds[i] < sorteds[j] })
	for i := range s {
		backtrack(sorteds, nil, i, &res)
	}
	return res
}

func backtrack(s, p []byte, i int, res *[]string) {
	if i < 0 || i >= len(s) || s[i] == '0' || (i > 0 && s[i] == s[i-1]) {
		return
	}
	p = append(p, s[i])
	if len(p) == len(s) {
		*res = append(*res, string(p))
		return
	}
	s[i] = '0'
	for j := range s {
		backtrack(s, p, j, res)
	}
	s[i] = p[len(p)-1]
}

func TestPermutation(t *testing.T) {
	t.Log(permutation("aab"))
}
