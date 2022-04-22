package _83_move_zeroes

func moveZeroes(nums []int) {
	index := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] != 0 {
			nums[index] = nums[i]
			index++
		}
	}
	for i := index; i < n; i++ {
		nums[i] = 0
	}
}

//func moveZeroes(nums []int) {
//	left, right, n := 0, 0, len(nums)
//	for right < n {
//		if nums[right] != 0 {
//			nums[right], nums[left] = nums[left], nums[right]
//			left++
//		}
//		right++
//	}
//}
