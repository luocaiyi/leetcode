package _32_implement_queue_using_stacks

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	queue := MyQueue{}
	queue.Push(1)
	queue.Push(2)
	fmt.Printf("%d; %d; %v", queue.Peek(), queue.Pop(), queue.Empty())
}
