package _42_linked_list_cycle_ii

import "github.com/luocaiyi/leetcode/datastruct"

type ListNode = datastruct.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 快慢指针
func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	for nil != fast && nil != fast.Next {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
}

// map
//func detectCycle(head *ListNode) *ListNode {
//	seen := map[*ListNode]struct{}{}
//	for nil != head {
//		if _, ok := seen[head]; ok {
//			return head
//		}
//		seen[head] = struct{}{}
//		head = head.Next
//	}
//	return nil
//}
