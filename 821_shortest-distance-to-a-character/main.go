package _21_shortest_distance_to_a_character

func shortestToChar(s string, c byte) []int {
	n := len(s)
	ans := make([]int, n)
	idx := -n
	for i, ch := range s {
		if byte(ch) == c {
			idx = i
		}
		ans[i] = i - idx
	}
	idx = n * 2
	for i := n - 1; i >= 0; i-- {
		if s[i] == c {
			idx = i
		}
		ans[i] = min(ans[i], idx-i)
	}
	return ans
}

//func shortestToChar(s string, c byte) []int {
//	var ans []int
//	for i := 0; i < len(s); i++ {
//		ans = append(ans, math.MaxInt)
//	}
//	for i := 0; i < len(s); i++ {
//		if c == s[i] {
//			ans[i] = 0
//			ans = giveValue(ans, i)
//		}
//	}
//	return ans
//}
//
//func giveValue(ans []int, index int) []int {
//	for i := 0; i < len(ans); i++ {
//		ans[i] = min(abs(index - i), ans[i])
//	}
//	return ans
//}
//
//func abs(x int) int {
//	if x < 0 {
//		return -x
//	}
//	return x
//}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}
