package _96_rotate_function

func maxRotateFunction(nums []int) int {
	sum, n, fn := 0, len(nums), 0
	for i := 0; i < n; i++ {
		sum += nums[i]
		fn += i * nums[i]
	}
	fnMax := fn
	for i := 1; i < n; i++ {
		fn = fn - sum + n*nums[i-1]
		fnMax = max(fnMax, fn)
	}
	return fnMax
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
