package _35_lowest_common_ancestor_of_a_binary_search_tree

import "github.com/luocaiyi/leetcode/datastruct"

type TreeNode = datastruct.TreeNode

// 递归
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return root
}

// 非递归
//func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
//	for root != nil {
//		if p.Val < root.Val && q.Val < root.Val {
//			root = root.Left
//		} else if p.Val > root.Val && q.Val > root.Val {
//			root = root.Right
//		} else {
//			return root
//		}
//	}
//	return nil
//}
