package _061_number_of_ways_to_buy_pens_and_pencils

import "testing"

func TestWaysToBuyPensPencils(t *testing.T) {
	total, cost1, cost2 := 20, 10, 5
	// 9
	println(waysToBuyPensPencils(total, cost1, cost2))
}
