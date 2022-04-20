package _523_count_odd_numbers_in_an_interval_range

func countOdds(low int, high int) int {
	return pre(high) - pre(low-1)
}

func pre(x int) int {
	return (x + 1) >> 1
}

//func countOdds(low int, high int) int {
//	if isEven(low) { low++ }
//	if isEven(high) { high-- }
//	return (high-low)/2 + 1
//}
//
//func isEven(num int) bool {
//	if num % 2 == 0 { return true }
//	return false
//}
