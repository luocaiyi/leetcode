package _243_calculate_digit_sum_of_a_string

import "testing"

func TestTest(t *testing.T) {
	println(digitSum("1234", 2))                   // 37
	println(digitSum("11111222223", 3))            // 135
	println(digitSum("11111", 4))                  // 41
	println(digitSum("00000000", 3))               // 000
	println(digitSum("742190887044879874008", 20)) // 978
}
