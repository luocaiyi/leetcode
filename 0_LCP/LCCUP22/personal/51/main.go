package _1

import "math/bits"

func perfectMenu(materials []int, cookbooks [][]int, attribute [][]int, limit int) (ans int) {
	calc := func(sub int) int {
		var cx, cy int
		cm := make([]int, len(materials))
		for _s := uint(sub); _s > 0; _s &= _s - 1 {
			p1 := bits.TrailingZeros(_s)
			a := cookbooks[p1]
			for i, v := range a {
				cm[i] += v
				if cm[i] > materials[i] {
					return -1
				}
			}
			cx += attribute[p1][0]
			cy += attribute[p1][1]
		}
		if cy < limit {
			return -1
		}
		return cx
	}
	ans = -1
	for sub := 0; sub < 1<<len(cookbooks); sub++ {
		res := calc(sub)
		ans = max(ans, res)
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
