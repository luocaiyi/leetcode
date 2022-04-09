package _swap_nodes_in_pairs

import "testing"

var head *ListNode

func init() {
	n5 := &ListNode{Val: 5}
	n4 := &ListNode{Val: 4, Next: n5}
	n3 := &ListNode{Val: 3, Next: n4}
	n2 := &ListNode{Val: 2, Next: n3}
	head = &ListNode{Val: 1, Next: n2}
}

func TestSwapPairs(t *testing.T) {
	head.Print()
	head = swapPairs(head)
	head.Print()
}
