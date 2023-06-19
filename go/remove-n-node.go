package main
import (
	"fmt"
)

func removeNInit() {
	head := &ListNodes{1, &ListNodes{2, &ListNodes{3, &ListNodes{4, &ListNodes{5, nil}}}}}
	n := 2
	newHead := removeNthFromEnd(head, n)

	// print out the resulting linked list
	for newHead != nil {
		fmt.Print(newHead.Val)
		newHead = newHead.Next
		if newHead != nil {
			fmt.Print(" -> ")
		}
	}
}

type ListNodes struct {
	Val int
	Next *ListNodes
}

func removeNthFromEnd(head *ListNodes, n int) *ListNodes {
	// create two pointers that are initially at the head of the list
	slow, fast := head, head

	// move the fast pointer n steps ahead
	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	// if the fast pointer reaches the end of the list, remove the head node
	if fast == nil {
		return head.Next
	}

	// move both pointers until the fast pointer reaches the end of the list
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	// remove the nth node by updating the next pointer of the (n-1)th node
	slow.Next = slow.Next.Next

	return head
}