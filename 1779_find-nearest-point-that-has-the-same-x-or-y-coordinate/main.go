package _779_find_nearest_point_that_has_the_same_x_or_y_coordinate

import "math"

func nearestValidPoint(x int, y int, points [][]int) int {
	ans := -1
	m := math.MaxInt
	for i, point := range points {
		if point[0] == x || point[1] == y {
			ab := abs(x-point[0]) + abs(y-point[1])
			if ab < m {
				m = ab
				ans = i
			}
		}
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
