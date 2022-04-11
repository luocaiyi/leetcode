package _03_kth_largest_element_in_a_stream

import (
	"container/heap"
	"sort"
)

type KthLargest struct {
	sort.IntSlice
	k int
}

func Constructor(k int, nums []int) KthLargest {
	k1 := KthLargest{k: k}
	for _, val := range nums {
		k1.Add(val)
	}
	return k1
}

func (k1 *KthLargest) Push(v interface{}) {
	k1.IntSlice = append(k1.IntSlice, v.(int))
}

func (k1 *KthLargest) Pop() interface{} {
	a := k1.IntSlice
	v := a[len(a)-1]
	k1.IntSlice = a[:len(a)-1]
	return v
}

func (k1 *KthLargest) Add(val int) int {
	heap.Push(k1, val)
	if k1.Len() > k1.k {
		heap.Pop(k1)
	}
	return k1.IntSlice[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
