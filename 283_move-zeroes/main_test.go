package _83_move_zeroes

import (
	"fmt"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	nums := []int{2, 1}
	moveZeroes(nums)
	fmt.Printf("%+v\n", nums)
}
