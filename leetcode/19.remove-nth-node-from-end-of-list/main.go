package main

type ListNode struct {
	Val int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var peek, cur = head, head

	for n > 0 {
		if peek.Next == nil {
			// assume n is valid, so n == len of List, remove head
			return head.Next
		}
		peek = peek.Next
		n --
	}
	for peek.Next != nil {
		peek = peek.Next
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	return head
}
