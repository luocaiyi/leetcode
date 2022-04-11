package __two_sum

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	res := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Printf("%v\n", res)
	res = twoSum([]int{3, 2, 4}, 6)
	fmt.Printf("%v\n", res)
}
