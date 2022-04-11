package _39_sliding_window_maximum

func maxSlidingWindow(nums []int, k int) []int {
	if nil == nums || len(nums) == 0 {
		return []int{}
	}
	var window, res []int
	for i, x := range nums {
		if i >= k && window[0] <= i-k {
			window = window[1:]
		}
		for len(window) != 0 && nums[window[len(window)-1]] <= x {
			window = window[:len(window)-1]
		}
		window = append(window, i)
		if i >= k-1 {
			res = append(res, nums[window[0]])
		}
	}
	return res
}
