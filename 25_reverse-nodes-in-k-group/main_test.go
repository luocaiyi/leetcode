package _5_reverse_nodes_in_k_group

import "testing"

var list, listsp *ListNode

func init() {
	n5 := &ListNode{Val: 5}
	n4 := &ListNode{Val: 4, Next: n5}
	n3 := &ListNode{Val: 3, Next: n4}
	n2 := &ListNode{Val: 2, Next: n3}
	n1 := &ListNode{Val: 1, Next: n2}
	list = n1

	listsp = &ListNode{Val: 1}
}

func TestReverseKGroup(t *testing.T) {
	list.Print()
	list = reverseKGroup(list, 2)
	list.Print()

	list = reverseKGroup(list, 2)
	list.Print()
	list = reverseKGroup(list, 3)
	list.Print()

	list = reverseKGroup(list, 3)
	list.Print()
	list = reverseKGroup(list, 1)
	list.Print()

	listsp.Print()
	listsp = reverseKGroup(listsp, 1)
	listsp.Print()
}
