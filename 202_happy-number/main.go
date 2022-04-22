package _02_happy_number

func isHappy(n int) bool {
	has := make(map[int]bool)
	has[n] = true
	for sum := 0; n != 1; n = sum {
		sum = 0
		for x := 0; n != 0; n /= 10 {
			x = n % 10
			sum += x * x
		}
		if sum == 1 {
			return true
		}
		if has[sum] {
			return false
		} else {
			has[sum] = true
		}
	}
	return true
}
