package datastruct

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (root *TreeNode) PreOrder() []int {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	var stack []*TreeNode
	var res []int
	stack = append(stack, root)
	for len(stack) != 0 {
		e := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, e.Val)
		if e.Right != nil {
			stack = append(stack, e.Right)
		}
		if e.Left != nil {
			stack = append(stack, e.Left)
		}
	}
	return res
}

func (root *TreeNode) InOrder() []int {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	res := root.Left.InOrder()
	fmt.Printf("-1:%v\n", res)
	res = append(res, root.Val)
	fmt.Printf("-2:%v\n", res)
	res = append(res, root.Right.InOrder()...)
	fmt.Printf("-3:%v\n", res)
	return res
}

func (root *TreeNode) PostOrder() []int {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	var res []int
	if root.Left != nil {
		lres := root.Left.PostOrder()
		if len(lres) > 0 {
			res = append(res, lres...)
		}
	}
	if root.Right != nil {
		rres := root.Right.PostOrder()
		if len(rres) > 0 {
			res = append(res, rres...)
		}
	}
	res = append(res, root.Val)
	return res
}
