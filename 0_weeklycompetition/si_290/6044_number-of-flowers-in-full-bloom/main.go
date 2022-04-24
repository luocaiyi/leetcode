package _044_number_of_flowers_in_full_bloom

import "sort"

// 超赞的 solution
// 正在开的花的个数 为所有在此时间节点及之前开放的花的个数减去在时间节点之前凋谢的花的个数
func fullBloomFlowers(flowers [][]int, persons []int) []int {
	n := len(flowers)
	op, cl := make([]int, n), make([]int, n)
	for i, flower := range flowers {
		op[i] = flower[0]
		cl[i] = flower[1]
	}
	sort.Ints(op)
	sort.Ints(cl)
	ans := make([]int, len(persons))
	for i, person := range persons {
		opNum := sort.SearchInts(op, person+1)
		clNum := sort.SearchInts(cl, person)
		ans[i] = opNum - clNum
	}
	return ans
}

// 差分
//func fullBloomFlowers(flowers [][]int, persons []int) []int {
//	diff := map[int]int{}
//	for _, flower := range flowers {
//		diff[flower[0]]++
//		diff[flower[1] + 1]--
//	}
//	n := len(diff)
//	times := make([]int, 0, n)
//	for t := range diff {
//		times = append(times, t)
//	}
//	sort.Ints(times)
//	for i, person := range persons {
//		persons[i] = person<<32 | i
//	}
//	sort.Ints(persons)
//
//	ans := make([]int, len(persons))
//	i, sum := 0, 0
//	for _, p := range persons {
//		for ; i < n && times[i] <= p>>32; i++ {
//			sum += diff[times[i]]
//		}
//		ans[uint32(p)] = sum
//	}
//	return ans
//}
