package _822_sign_of_the_product_of_an_array

func arraySign(nums []int) int {
	cnt := 0
	for _, num := range nums {
		if num == 0 {
			return 0
		}
		if num < 0 {
			cnt++
		}
	}
	if cnt%2 == 0 {
		return 1
	}
	return -1
}
