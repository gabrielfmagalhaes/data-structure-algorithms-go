package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

var visited []*ListNode

func main() {
	nodeInit := &ListNode{Val: 10}
	nodeMiddle := &ListNode{Val: 2}
	nodeEnd := &ListNode{Val: 5}

	nodeInit.Next = nodeMiddle
	nodeMiddle.Next = nodeEnd
	nodeEnd.Next = nodeInit

	result := hasCycle(nodeInit)

	fmt.Printf("Result is %v", result)
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	fast := head
	slow := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if slow == fast {
			return true
		}
	}

	return false
}
