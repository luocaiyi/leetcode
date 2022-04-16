package _85_mini_parser

import (
	"strconv"
	"unicode"
)

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (n NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (n NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (n *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (n NestedInteger) GetList() []*NestedInteger {}
 */
func deserialize(s string) *NestedInteger {
	if s[0] != '[' {
		num, _ := strconv.Atoi(s)
		nested := &NestedInteger{}
		nested.SetInteger(num)
		return nested
	}
	var stack []*NestedInteger
	num, negative := 0, false
	for i, ch := range s {
		if ch == '-' {
			negative = true
		} else if unicode.IsDigit(ch) {
			num = num*10 + int(ch-'0')
		} else if ch == '[' {
			stack = append(stack, &NestedInteger{})
		} else if ch == ',' || ch == ']' {
			if unicode.IsDigit(rune(s[i-1])) {
				if negative {
					num = -num
				}
				nested := NestedInteger{}
				nested.SetInteger(num)
				stack[len(stack)-1].Add(nested)
			}
			num, negative = 0, false
			if ch == ']' && len(stack) > 1 {
				stack[len(stack)-2].Add(*stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
		}
	}
	return stack[len(stack)-1]
}

// 递归 dfs() 深度优先搜索
//func deserialize(s string) *NestedInteger {
//	index := 0
//	var dfs func() *NestedInteger
//	dfs = func() *NestedInteger {
//		nested := &NestedInteger{}
//		if s[index] == '[' {
//			index++
//			for s[index] != ']' {
//				nested.Add(*dfs())
//				if s[index] == ',' {
//					index++
//				}
//			}
//			index++
//			return nested
//		}
//		negative := s[index] == '-'
//		if negative {
//			index++
//		}
//		num := 0
//		for ; index < len(s) && unicode.IsDigit(rune(s[index])) ; index++ {
//			num = num * 10 + int(s[index]-'0')
//		}
//		if negative {
//			num = -num
//		}
//		nested.SetInteger(num)
//		return nested
//	}
//	return dfs()
//}
