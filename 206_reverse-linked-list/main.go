package _reverse_linked_list

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// Print 打印链表
func (head *ListNode) Print() {
	cur := head
	format := "["
	for nil != cur {
		format += fmt.Sprintf("%d", cur.Val)
		cur = cur.Next
		if nil != cur {
			format += ","
		}
	}
	format += "]"
	fmt.Println(format)
}

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
