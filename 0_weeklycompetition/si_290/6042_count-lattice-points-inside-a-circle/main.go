package _042_count_lattice_points_inside_a_circle

func countLatticePoints(circles [][]int) int {
	m := map[[2]int]struct{}{}
	for _, circle := range circles {
		x, y, r := circle[0], circle[1], circle[2]
		m[[2]int{x, y}] = struct{}{}
		m[[2]int{x + r, y}] = struct{}{}
		m[[2]int{x - r, y}] = struct{}{}
		m[[2]int{x, y + r}] = struct{}{}
		m[[2]int{x, y - r}] = struct{}{}
		for i := 1; i < r; i++ {
			for j := -i; j <= i; j++ {
				if inCircle(i, j, r) {
					m[[2]int{x + i, y + j}] = struct{}{}
					m[[2]int{x - i, y + j}] = struct{}{}
					m[[2]int{x + j, y + i}] = struct{}{}
					m[[2]int{x + j, y - i}] = struct{}{}
				}
			}
		}
	}
	return len(m)
}

func inCircle(x, y, r int) bool {
	if x*x+y*y <= r*r {
		return true
	}
	return false
}
