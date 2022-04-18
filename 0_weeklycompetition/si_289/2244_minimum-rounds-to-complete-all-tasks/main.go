package _244_minimum_rounds_to_complete_all_tasks

func minimumRounds(tasks []int) int {
	count := map[int]int{}
	for _, k := range tasks {
		count[k]++
	}
	ans := 0
	for _, v := range count {
		if v == 1 {
			return -1
		}
		ans += (v + 2) / 3
	}
	return ans
}
