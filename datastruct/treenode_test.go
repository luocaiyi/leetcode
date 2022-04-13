package datastruct

import (
	"fmt"
	"testing"
)

var root *TreeNode

func init() {
	root = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{
			Val:   5,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 7},
		},
	}
}

func TestTreeNode_PreOrder(t *testing.T) {
	res := root.PreOrder()
	fmt.Printf("%v\n", res)
}

func TestTreeNode_InOrder(t *testing.T) {
	res := root.InOrder()
	fmt.Printf("%v\n", res)
}

func TestTreeNode_PostOrder(t *testing.T) {
	res := root.PostOrder()
	fmt.Printf("%v\n", res)
}
