package _041_intersection_of_multiple_arrays

import "sort"

func intersection(nums [][]int) (ans []int) {
	m := map[int]int{}
	for _, num := range nums {
		for _, n := range num {
			m[n]++
		}
	}
	n := len(nums)
	for k, v := range m {
		if v == n {
			ans = append(ans, k)
		}
	}
	sort.Ints(ans)
	return
}
