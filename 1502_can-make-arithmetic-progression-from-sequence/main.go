package _502_can_make_arithmetic_progression_from_sequence

import "sort"

func canMakeArithmeticProgression(arr []int) bool {
	if len(arr) == 2 {
		return true
	}
	sort.Ints(arr)
	x := arr[1] - arr[0]
	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] != x {
			return false
		}
	}
	return true
}
