package _0_valid_parentheses

func isValid(s string) bool {
	l := len(s)
	if l < 0 || l%2 != 0 {
		return false
	}

	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []byte{}
	for i := 0; i < l; i++ {
		if c, ok := pairs[s[i]]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != c {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}

// ASCII取巧，实用需优化
//func isValid(s string) bool {
//	l := len(s)
//	if l < 0 || l % 2 != 0 {
//		return false
//	}
//
//	stack := []byte{}
//	for i := 0; i < l; i++ {
//		if len(stack) == 0 {
//			stack = append(stack, s[i])
//		} else if s[i]-stack[len(stack)-1] == 1 || s[i]-stack[len(stack)-1] == 2 {
//			stack = stack[:len(stack)-1]
//		} else {
//			stack = append(stack, s[i])
//		}
//	}
//	return len(stack) == 0
//}
