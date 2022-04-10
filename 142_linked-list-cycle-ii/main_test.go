package _42_linked_list_cycle_ii

import (
	"log"
	"testing"
)

var headCycle1, resCycle1, headCycle2, resCycle2, headNormal *ListNode

func init() {
	n4 := &ListNode{Val: -4}
	n3 := &ListNode{Val: 0, Next: n4}
	n2 := &ListNode{Val: 2, Next: n3}
	n1 := &ListNode{Val: 3, Next: n2}
	headCycle1 = n1
	n4.Next = n2
	resCycle1 = n2

	n6 := &ListNode{Val: 2}
	n5 := &ListNode{Val: 1, Next: n6}
	headCycle2 = n5
	resCycle2 = n5
	n6.Next = n5

	n7 := &ListNode{Val: 1}
	headNormal = n7
}

func TestDetectCycle(t *testing.T) {
	res1 := resCycle1 == detectCycle(headCycle1)
	res2 := resCycle2 == detectCycle(headCycle2)
	res3 := nil == detectCycle(headNormal)
	if !res1 || !res2 || !res3 {
		log.Fatal("Err")
	}
}
