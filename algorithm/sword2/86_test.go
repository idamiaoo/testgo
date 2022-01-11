package sword2

func partition(s string) [][]string {
	var ans [][]string
	var path []string
	var backtracking func(index int)
	backtracking = func(index int) {
		if index == len(s) {
			ans = append(ans, append([]string{}, path...))
			return
		}
		for i := index; i < len(s); i++ {
			if isPalindrome2(s, index, i) {
				path = append(path, s[index:i+1])
				backtracking(i + 1)
				path = path[:len(path)-1]
			}
		}
	}
	backtracking(0)
	return ans
}

func isPalindrome2(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
