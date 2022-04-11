package _03_kth_largest_element_in_a_stream

import (
	"testing"
)

func TestKthLargest(t *testing.T) {
	k1 := Constructor(3, []int{4, 5, 8, 2})
	k := k1.Add(3)
	println(k)
	k = k1.Add(5)
	println(k)
	k = k1.Add(10)
	println(k)
	k = k1.Add(9)
	println(k)
	k = k1.Add(4)
	println(k)
}
