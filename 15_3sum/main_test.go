package _5_3sum

import (
	"fmt"
	"testing"
)

func TestThreeSum(t *testing.T) {
	res := threeSum([]int{-1, 0, 1, 2, -1, -4})
	fmt.Printf("%v\n", res)
	res = threeSum([]int{})
	fmt.Printf("%v\n", res)
	res = threeSum([]int{0})
	fmt.Printf("%v\n", res)
	res = threeSum([]int{0, 0, 0, 0, 0})
	fmt.Printf("%v\n", res)
	res = threeSum([]int{-2, 0, 1, 1, 2})
	fmt.Printf("%v\n", res)
}
