package _linked_list_cycle

import datastruct "github.com/luocaiyi/leetcode/datastruct"

type ListNode = datastruct.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 快慢指针
func hasCycle(head *ListNode) bool {
	fast, slow := head, head
	for nil != fast && nil != slow && nil != fast.Next {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

// map
//func hasCycle(head *ListNode) bool {
//	seen := map[*ListNode]struct{}{}
//	for nil != head {
//		if _, ok := seen[head]; ok {
//			return true
//		}
//		seen[head] = struct{}{}
//		head = head.Next
//	}
//	return false
//}
