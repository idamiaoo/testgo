package sword2

func pow(a, b int) int {
	if b == 0 {
		return 1
	}
	if b == 1 {
		return a
	}
	return pow(a*a, b/2) * pow(a, b%2)
}
