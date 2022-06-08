package leetcode

func removeAnagrams(words []string) []string {
	if len(words) == 1 {
		return words
	}
	var removeIndex = make(map[int]bool)
	for i, j := 0, 1; j < len(words); j++ {
		if isAnagram(words[i], words[j]) {
			removeIndex[j] = true
		} else {
			i = j
		}
	}

	if len(removeIndex) == 0 {
		return words
	}
	var ans []string
	for i := range words {
		if removeIndex[i] {
			continue
		}
		ans = append(ans, words[i])
	}
	return ans
}

func isAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	m1, m2 := make(map[byte]int), make(map[byte]int)
	for _, v := range []byte(s1) {
		m1[v]++
	}
	for _, v := range []byte(s2) {
		m2[v]++
	}
	for k, v := range m1 {
		if m2[k] != v {
			return false
		}
	}
	return true
}
