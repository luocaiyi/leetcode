package _66_reshape_the_matrix

import (
	"fmt"
	"testing"
)

func TestMatrixReshape(t *testing.T) {
	mat := [][]int{{1, 2}, {3, 4}}
	fmt.Printf("%+v\n", matrixReshape(mat, 1, 4))
	fmt.Printf("%+v\n", matrixReshape(mat, 4, 1))
}
