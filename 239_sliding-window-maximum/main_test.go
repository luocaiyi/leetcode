package _39_sliding_window_maximum

import (
	"fmt"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	res := maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)
	fmt.Printf("%v", res)
}
