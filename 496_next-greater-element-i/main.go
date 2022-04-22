package _96_next_greater_element_i

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	mp := map[int]int{}
	var stack []int
	for i := len(nums2) - 1; i >= 0; i-- {
		num := nums2[i]
		for len(stack) > 0 && num >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			mp[num] = stack[len(stack)-1]
		} else {
			mp[num] = -1
		}
		stack = append(stack, num)
	}
	ans := make([]int, len(nums1))
	for i, num := range nums1 {
		ans[i] = mp[num]
	}
	return ans
}
