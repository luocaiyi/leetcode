package _25_implement_stack_using_queues

import (
	"fmt"
	"testing"
)

func TestMyStack(t *testing.T) {
	myStack := &MyStack{}
	myStack.Push(1)
	myStack.Push(2)
	fmt.Printf("%d; %d; %v", myStack.Top(), myStack.Pop(), myStack.Empty())
}
