package datastruct

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
