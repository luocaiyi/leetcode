package _89_n_ary_tree_preorder_traversal

import "github.com/luocaiyi/leetcode/datastruct"

type Node = datastruct.Node

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) (ans []int) {
	var dfs func(node *Node)
	dfs = func(node *Node) {
		if node == nil {
			return
		}
		ans = append(ans, node.Val)
		for _, child := range node.Children {
			dfs(child)
		}
	}
	dfs(root)
	return
}

//func preorder(root *Node) (ans []int) {
//	if root == nil {
//		return nil
//	}
//	ans = append(ans, root.Val)
//	for _, child := range root.Children {
//		ans = append(ans, preorder(child)...)
//	}
//	return
//}
