package main

import (
    "container/heap"
    "fmt"
)

// Definition for singly-linked list.
type MergeKListNode struct {
    Val int
    Next *MergeKListNode
}

// PriorityQueue implements heap.Interface and holds MergeKListNodes.
type PriorityQueue []*MergeKListNode

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
    return pq[i].Val < pq[j].Val
}

func (pq PriorityQueue) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
    item := x.(*MergeKListNode)
    *pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
    old := *pq
    n := len(old)
    item := old[n-1]
    *pq = old[0 : n-1]
    return item
}

func mergeKLists(lists []*MergeKListNode) *MergeKListNode {
    // Initialize PriorityQueue and push all list heads
    pq := make(PriorityQueue, 0, len(lists))
    for i := 0; i < len(lists); i++ {
        if lists[i] != nil {
            pq = append(pq, lists[i])
        }
    }
    heap.Init(&pq)

    // Initialize dummy node for result list
    dummy := &MergeKListNode{}
    curr := dummy

    // Pop the smallest node from the PriorityQueue and push its next node
    for pq.Len() > 0 {
        smallestNode := heap.Pop(&pq).(*MergeKListNode)
        curr.Next = smallestNode
        curr = curr.Next
        if smallestNode.Next != nil {
            heap.Push(&pq, smallestNode.Next)
        }
    }

    return dummy.Next
}

func mergeKListsInit() {
    // Create test case and print result
    n1 := &MergeKListNode{Val: 1, Next: &MergeKListNode{Val: 4, Next: &MergeKListNode{Val: 5}}}
    n2 := &MergeKListNode{Val: 1, Next: &MergeKListNode{Val: 3, Next: &MergeKListNode{Val: 4}}}
    n3 := &MergeKListNode{Val: 2, Next: &MergeKListNode{Val: 6}}
    lists := []*MergeKListNode{n1, n2, n3}

    result := mergeKLists(lists)
    for result != nil {
        fmt.Print(result.Val, " ")
        result = result.Next
	}
}
