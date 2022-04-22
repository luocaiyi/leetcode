package _672_richest_customer_wealth

func maximumWealth(accounts [][]int) (ans int) {
	for _, account := range accounts {
		sum := 0
		for _, m := range account {
			sum += m
		}
		ans = max(ans, sum)
	}
	return
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
