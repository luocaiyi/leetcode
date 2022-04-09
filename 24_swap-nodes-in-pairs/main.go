package _swap_nodes_in_pairs

import datastruct "github.com/luocaiyi/leetcode/datastruct"

type ListNode = datastruct.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 迭代
//func swapPairs(head *ListNode) *ListNode {
//	if nil == head || nil == head.Next {
//		return head
//	}
//
//	newHead := head.Next
//	prev := &ListNode{Next: head}
//	for nil != prev.Next && nil != prev.Next.Next {
//		//prev, prev.Next, prev.Next.Next, prev.Next.Next.Next = prev.Next, prev.Next.Next , prev.Next.Next.Next, prev.Next
//		a := prev.Next
//		b := prev.Next.Next
//		prev.Next, b.Next, a.Next = b, a , b.Next
//		prev = a
//	}
//	return newHead
//}

// 递归
func swapPairs(head *ListNode) *ListNode {
	if nil == head || nil == head.Next {
		return head
	}

	newHead := head.Next
	head.Next = swapPairs(newHead.Next)
	newHead.Next = head
	return newHead
}
