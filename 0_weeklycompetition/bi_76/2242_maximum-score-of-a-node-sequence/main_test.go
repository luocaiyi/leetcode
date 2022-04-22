package _242_maximum_score_of_a_node_sequence

import "testing"

func TestMaximumScore(t *testing.T) {
	scores := []int{5, 2, 9, 8, 4}
	edged := [][]int{{0, 1}, {1, 2}, {2, 3}, {0, 2}, {1, 3}, {2, 4}}
	// 24
	println(maximumScore(scores, edged))
	scores = []int{9, 20, 6, 4, 11, 12}
	edged = [][]int{{0, 3}, {5, 3}, {2, 4}, {1, 3}}
	// -1
	println(maximumScore(scores, edged))
}
