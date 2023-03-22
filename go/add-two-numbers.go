package main
import "fmt"
func twoNumbersinit() {
	l1 := &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 2, Next: nil}}}
	l2 := &ListNode{Val: 4, Next: &ListNode{Val: 6, Next: &ListNode{Val: 5, Next: nil}}}
	result := addTwoNumbers(l1, l2)
	for result != nil {
		fmt.Printf("%d | ", result.Val)
		result = result.Next
	}
}

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	carry := 0

	for l1 != nil || l2 != nil {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		curr.Next = &ListNode{Val: sum % 10}
		curr = curr.Next
		carry = sum / 10
	}

	if carry > 0 {
		curr.Next = &ListNode{Val: carry}
	}

	return dummy.Next
}