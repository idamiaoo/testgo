package leetcode

func rearrangeCharacters(s string, target string) int {
	m := make(map[byte]int)
	for _, v := range []byte(s) {
		m[v]++
	}
	var (
		ans  int
		exit bool
	)
	for !exit {
		for _, v := range []byte(target) {
			if m[v] == 0 {
				exit = true
				break
			} else {
				m[v]--
			}
		}
		if !exit {
			ans++
		}
	}
	return ans
}
