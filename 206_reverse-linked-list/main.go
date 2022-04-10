package _reverse_linked_list

import "github.com/luocaiyi/leetcode/datastruct"

type ListNode = datastruct.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func ReverseList(head *ListNode) *ListNode {
	if nil == head || nil == head.Next {
		return head
	}

	var prev *ListNode = nil
	cur := head
	for nil != cur {
		cur.Next, prev, cur = prev, cur, cur.Next
	}
	return prev
}
