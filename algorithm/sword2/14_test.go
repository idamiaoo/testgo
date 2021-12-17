package sword2

import (
	"testing"
)

func checkInclusion(s1 string, s2 string) bool {
	m1 := make(map[byte]int)
	for _, v := range []byte(s1) {
		m1[v]++
	}
	m := make(map[byte]int)

	for i, j := 0, 0; j < len(s2); {
		if _, ok := m1[s2[j]]; ok {
			m[s2[j]]++
			if j-i+1 == len(s1) {
				var inclusion = true
				for k, v := range m1 {
					if m[k] != v {
						inclusion = false
					}
				}
				if inclusion {
					return true
				}
				m[s2[i]]--
				i++
			}
			j++
		} else {
			m = make(map[byte]int)
			j++
			i = j
		}

	}
	return false
}

func Test14(t *testing.T) {
	var s1, s2 = "ab", "eidbaooo"
	var s3, s4 = "adc", "dcda"

	t.Log(checkInclusion(s1, s2))
	t.Log(checkInclusion(s3, s4))

}
