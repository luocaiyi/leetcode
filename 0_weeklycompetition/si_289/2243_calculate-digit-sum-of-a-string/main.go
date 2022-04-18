package _243_calculate_digit_sum_of_a_string

import (
	"fmt"
	"strconv"
)

// 递归
func digitSum(s string, k int) string {
	if len(s) <= k {
		return s
	}
	ans := ""
	for i := 0; i < len(s); i += k {
		r := 0
		for _, v := range s[i:min(i+k, len(s))] {
			n, _ := strconv.Atoi(string(v))
			r += n
		}
		ans += fmt.Sprintf("%d", r)
	}
	return digitSum(ans, k)
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

// 迭代
//func digitSum(s string, k int) string {
//	ans := s
//	for len(s) > k {
//		ans = ""
//		for {
//			sum := 0
//			for i := 0; i < len(s) && i < k; i++ {
//				n, _ := strconv.Atoi(string(s[i]))
//				sum += n
//			}
//			ans += fmt.Sprintf("%d", sum)
//			if len(s) <= k {
//				break
//			}
//			s = s[k:]
//		}
//		s = ans
//	}
//	return ans
//}
