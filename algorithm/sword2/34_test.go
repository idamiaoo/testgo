package sword2

func isAlienSorted(words []string, order string) bool {
	var orderArr [26]int
	for i, v := range []byte(order) {
		orderArr[v-'a'] = i + 1
	}
	var compare = func(a, b string) int {
		n := len(a)
		if len(b) > len(a) {
			n = len(b)
		}
		for i := 0; i < n; i++ {
			if i == len(a) {
				return -1
			}
			if i == len(b) {
				return 1
			}
			v := orderArr[a[i]-'a'] - orderArr[b[i]-'a']
			if v < 0 {
				return -1
			}
			if v > 0 {
				return 1
			}
		}
		return 0
	}
	if len(words) == 1 {
		return true
	}

	for i := 0; i < len(words)-1; i++ {
		if compare(words[i], words[i+1]) > 0 {
			return false
		}
	}
	return true
}
