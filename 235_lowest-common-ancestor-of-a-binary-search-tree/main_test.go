package _35_lowest_common_ancestor_of_a_binary_search_tree

import "testing"

func TestLowestCommonAncestor(t *testing.T) {
	p := &TreeNode{Val: 5}
	q := &TreeNode{Val: 9}
	tree := &TreeNode{
		Val: 6,
		Left: &TreeNode{
			Val:  2,
			Left: &TreeNode{Val: 0},
			Right: &TreeNode{
				Val:   4,
				Left:  &TreeNode{Val: 3},
				Right: q,
			},
		},
		Right: &TreeNode{
			Val:   8,
			Left:  &TreeNode{Val: 7},
			Right: q,
		},
	}

	res := lowestCommonAncestor(tree, p, q)
	println(res.Val)
}
