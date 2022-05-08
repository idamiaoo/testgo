package leetcode

func s71(n int) int {
	if n <= 2 {
		return n
	}
	a, b := 1, 2
	var c int
	for i := 3; i <= n; i++ {
		c = a + b
		a = b
		b = c
	}
	return c
}
