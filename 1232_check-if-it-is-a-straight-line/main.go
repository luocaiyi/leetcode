package _232_check_if_it_is_a_straight_line

func checkStraightLine(coordinates [][]int) bool {
	zeroX, zeroY := coordinates[0][0], coordinates[0][1]
	for _, coordinate := range coordinates {
		coordinate[0] -= zeroX
		coordinate[1] -= zeroY
	}
	A, B := coordinates[1][1], -coordinates[1][0]
	for _, P := range coordinates[2:] {
		x, y := P[0], P[1]
		if x*A+y*B != 0 {
			return false
		}
	}
	return true
}
