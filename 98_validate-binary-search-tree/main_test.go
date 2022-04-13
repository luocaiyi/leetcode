package _8_validate_binary_search_tree

import (
	"fmt"
	"testing"
)

func TestIsValidBST(t *testing.T) {
	tree := &TreeNode{
		2,
		&TreeNode{1, nil, nil},
		&TreeNode{3, nil, nil},
	}

	fmt.Printf("%v\n", isValidBST(tree))
}
