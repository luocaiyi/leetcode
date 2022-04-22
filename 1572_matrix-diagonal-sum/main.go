package _572_matrix_diagonal_sum

func diagonalSum(mat [][]int) (ans int) {
	n := len(mat)
	for i, j := 0, n-1; i < n; i++ {
		if i == j {
			ans += mat[i][i]
		} else {
			ans += mat[i][i] + mat[i][j]
		}
		j--
	}
	return
}
