package _5_reverse_nodes_in_k_group

import "github.com/luocaiyi/leetcode/datastruct"

type ListNode = datastruct.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	hair := &ListNode{Next: head}
	pre := hair

	for nil != head {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if nil == tail {
				return hair.Next
			}
		}
		head, tail = reverseList(head, tail)
		pre.Next = head
		pre = tail
		head = tail.Next
	}
	return hair.Next
}

func reverseList(head, tail *ListNode) (*ListNode, *ListNode) {
	nex := tail.Next
	prev := nex
	cur := head
	for nex != cur {
		cur.Next, prev, cur = prev, cur, cur.Next
	}
	return tail, head
}
