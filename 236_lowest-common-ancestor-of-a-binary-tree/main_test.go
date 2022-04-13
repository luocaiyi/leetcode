package _36_lowest_common_ancestor_of_a_binary_tree

import "testing"

func TestLowestCommonAncestor(t *testing.T) {
	p := &TreeNode{Val: 6}
	q := &TreeNode{Val: 4}
	tree := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:  5,
			Left: p,
			Right: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 7},
				Right: q,
			},
		},
		Right: &TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: 0},
			Right: &TreeNode{Val: 8},
		},
	}
	res := lowestCommonAncestor(tree, p, q)
	println(res.Val)
}
