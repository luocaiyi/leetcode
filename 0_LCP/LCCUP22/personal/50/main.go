package _0

import (
	"math"
)

func giveGem(gem []int, operations [][]int) int {
	for _, op := range operations {
		num := gem[op[0]] / 2
		gem[op[0]] = gem[op[0]] - num
		gem[op[1]] = gem[op[1]] + num
	}
	mx := math.MinInt
	mi := math.MaxInt
	for _, v := range gem {
		if v > mx {
			mx = v
		}
		if v < mi {
			mi = v
		}
	}
	return mx - mi
}
