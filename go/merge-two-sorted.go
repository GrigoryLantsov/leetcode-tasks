package main
import (
	"fmt"
)

func mergeInit() {
	// create list1: 1 -> 2 -> 4
	l1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	// create list2: 1 -> 3 -> 4
	l2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}

	// merge the two lists
	merged := mergeTwoLists(l1, l2)

	// print the merged list
	curr := merged
	for curr != nil {
		fmt.Println(curr.Val)
		curr = curr.Next
	}
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// create a dummy head for the merged list
	dummy := &ListNode{}
	// create a pointer to the current node of the merged list
	curr := dummy

	// iterate through both lists while they are not empty
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			// if the value of the current node in list1 is less than that of list2, append it to the merged list and move to the next node in list1
			curr.Next = l1
			l1 = l1.Next
		} else {
			// otherwise, append the current node in list2 to the merged list and move to the next node in list2
			curr.Next = l2
			l2 = l2.Next
		}
		// move to the next node in the merged list
		curr = curr.Next
	}

	// append any remaining nodes in list1 or list2 to the merged list
	if l1 != nil {
		curr.Next = l1
	} else {
		curr.Next = l2
	}

	// return the head of the merged list (excluding the dummy head)
	return dummy.Next
}