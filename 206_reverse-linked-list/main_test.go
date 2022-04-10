package _reverse_linked_list

import "testing"

var head *ListNode

func init() {
	n5 := &ListNode{Val: 5}
	n4 := &ListNode{Val: 4, Next: n5}
	n3 := &ListNode{Val: 3, Next: n4}
	n2 := &ListNode{Val: 2, Next: n3}
	head = &ListNode{Val: 1, Next: n2}
}

func TestReverseList(t *testing.T) {
	head.Print()
	head = reverseList(head)
	head.Print()
}
