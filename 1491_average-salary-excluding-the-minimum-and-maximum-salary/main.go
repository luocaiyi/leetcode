package _491_average_salary_excluding_the_minimum_and_maximum_salary

func average(salary []int) float64 {
	av := float64(0)
	ma, mi := salary[0], salary[0]
	for _, v := range salary {
		ma = max(ma, v)
		mi = min(mi, v)
		av += float64(v)
	}
	av = av - float64(ma) - float64(mi)
	return av / float64(len(salary)-2)
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}
