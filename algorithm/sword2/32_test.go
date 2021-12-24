package sword2

func isAnagram(s string, t string) bool {
	if s == t || len(s) != len(t) {
		return false
	}
	runes := make(map[rune]int)
	for _, v := range s {
		runes[v] = runes[v] + 1
	}
	for _, v := range t {
		runes[v] = runes[v] - 1
	}
	for _, v := range runes {
		if v != 0 {
			return false
		}
	}
	return true
}
