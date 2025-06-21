package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */


type ListNode struct {
    Val int
    Next *ListNode
}

func hasCycle(head *ListNode) bool {
    slow := head
	fast := head
	for fast != nil && fast.Next != nil && slow != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}