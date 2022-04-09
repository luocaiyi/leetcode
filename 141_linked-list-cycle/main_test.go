package _linked_list_cycle

import (
	"log"
	"testing"
)

var headCycle, headNormal *ListNode

func init() {
	n5 := &ListNode{Val: 5}
	n4 := &ListNode{Val: 4, Next: n5}
	n3 := &ListNode{Val: 3, Next: n4}
	n2 := &ListNode{Val: 2, Next: n3}
	headCycle = &ListNode{Val: 1, Next: n2}
	n5.Next = n3
	n10 := &ListNode{Val: 5}
	n9 := &ListNode{Val: 4, Next: n10}
	n8 := &ListNode{Val: 3, Next: n9}
	n7 := &ListNode{Val: 2, Next: n8}
	headNormal = &ListNode{Val: 1, Next: n7}
}

func TestHasCycle(t *testing.T) {
	if hasCycle(headCycle) != true || hasCycle(headNormal) != false {
		log.Fatal("err")
	}
}
