package sword2

import (
	"testing"
)

func findAnagrams(s string, p string) []int {
	m1 := make(map[byte]int)
	for _, v := range []byte(p) {
		m1[v]++
	}
	m := make(map[byte]int)
	var res []int
	for i, j := 0, 0; j < len(s); {
		if _, ok := m1[s[j]]; ok {
			m[s[j]]++
			if j-i+1 == len(p) {
				var inclusion = true
				for k, v := range m1 {
					if m[k] != v {
						inclusion = false
					}
				}
				if inclusion {
					res = append(res, i)
				}
				m[s[i]]--
				i++
			}
			j++
		} else {
			m = make(map[byte]int)
			j++
			i = j
		}

	}
	return res
}

func Test15(t *testing.T) {
	var s1, s2 = "cbaebabacd", "abc"

	t.Log(findAnagrams(s1, s2))

}
