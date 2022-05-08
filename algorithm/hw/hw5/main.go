package main

import "fmt"

func atoi(s string) int {
	var (
		ans   int
		sign  int
		start int
	)
	for start = 0; start < len(s); start++ {
		if s[start] == ' ' {
			continue
		}
		break
	}
	sign = 1
	fmt.Println(start)
	for i := start; i < len(s); i++ {
		if s[i] == '-' && i == start {
			sign = -1
		} else if s[i] == '+' && i == start {
			sign = 1
		} else if s[i] >= '0' && s[i] <= '9' {
			ans = ans*10 + int(s[i]-'0')
		} else {
			return 0
		}
	}
	return ans * sign
}
