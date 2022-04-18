package template

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}
