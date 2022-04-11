package _8_4sum

import (
	"fmt"
	"testing"
)

func TestFourSum(t *testing.T) {
	res := fourSum([]int{1, 0, -1, 0, -2, 2}, 0)
	fmt.Printf("%v\n", res)
	res = fourSum([]int{2, 2, 2, 2, 2}, 8)
	fmt.Printf("%v\n", res)
	res = fourSum([]int{-2, -1, -1, 1, 1, 2, 2}, 0)
	fmt.Printf("%v\n", res)
}
