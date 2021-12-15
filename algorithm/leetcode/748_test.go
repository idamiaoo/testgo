package leetcode

func shortestCompletingWord(licensePlate string, words []string) string {
	counter := func(w string) ([]int, int) {
		cnts := make([]int, 26)
		var cnt int
		for _, v := range []byte(w) {
			if v >= 'a' && v <= 'z' {
				cnts[v-'a']++
				cnt++
			}
			if v >= 'A' && v <= 'Z' {
				cnts[v-'A']++
				cnt++
			}
		}
		return cnts, cnt
	}

	cnts, cnt := counter(licensePlate)

	var l int
	var ans string
	for _, w := range words {
		if (l == 0 || len(w) < l) && len(w) >= cnt {
			_cnts, _ := counter(w)
			var i int
			for ; i < 26; i++ {
				if _cnts[i] < cnts[i] {
					break
				}
			}
			if i == 26 {
				ans, l = w, len(w)
			}
		}
	}
	return ans
}
