package _1

import (
	"testing"
)

func TestPerfectMenu(t *testing.T) {
	materials := []int{3, 2, 4, 1, 2}
	cookbooks := [][]int{{1, 1, 0, 1, 2}, {2, 1, 4, 0, 0}, {3, 2, 4, 1, 0}}
	attribute := [][]int{{3, 2}, {2, 4}, {7, 6}}
	limit := 5
	// ans = 7
	println(perfectMenu(materials, cookbooks, attribute, limit))
}
