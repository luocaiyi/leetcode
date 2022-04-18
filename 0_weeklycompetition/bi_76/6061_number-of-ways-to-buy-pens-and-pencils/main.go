package _061_number_of_ways_to_buy_pens_and_pencils

func waysToBuyPensPencils(total int, cost1 int, cost2 int) int64 {
	var ans int64 = 0
	for i := 0; i <= total/cost1; i++ {
		ans += int64((total-i*cost1)/cost2 + 1)
	}
	return ans
}
