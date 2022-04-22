package _96_rotate_function

import "testing"

func TestMaxRotateFunction(t *testing.T) {
	nums := []int{4, 3, 2, 6}
	println(maxRotateFunction(nums))
	nums = []int{100}
	println(maxRotateFunction(nums))
}
