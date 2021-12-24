package sword2

import (
	"testing"
)

func minWindow(s string, t string) string {
	var m1, m2 = make(map[byte]int), make(map[byte]int)
	for _, v := range []byte(t) {
		m1[v]++
	}
	var cnt int
	var res string
	for i, j := 0, 0; j < len(s); j++ {
		m2[s[j]]++
		if m2[s[j]] <= m1[s[j]] {
			cnt++
		}
		for m2[s[i]] > m1[s[i]] && i < j {
			m2[s[i]]--
			i++
		}
		if cnt == len(t) {
			if res == "" || j-i+1 < len(res) {
				res = s[i : j+1]
			}
		}
	}
	return res
}

func Test17(t *testing.T) {
	var s, tt = "ADOBECODEBANC", "ABC"
	t.Log(minWindow(s, tt))
}
