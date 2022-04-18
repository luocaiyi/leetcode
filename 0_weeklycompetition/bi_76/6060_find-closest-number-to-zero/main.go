package _060_find_closest_number_to_zero

func findClosestNumber(nums []int) int {
	var ans = nums[0]
	for _, num := range nums {
		if abs(num) < abs(ans) || abs(num) == abs(ans) && num > ans {
			ans = num
		}
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
