package main

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sum := l1.Val + l2.Val
	overflow := sum / 10
	root := &ListNode{sum % 10, nil}
	var pre = root
	var cur *ListNode
	l1, l2 = l1.Next, l2.Next
	for l1 != nil || l2 != nil || overflow > 0 {
		var x, y int
		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}

		sum = x + y + overflow
		overflow = sum / 10
		cur = &ListNode{sum % 10, nil}
		pre.Next = cur
		pre = cur
		continue
	}
	return root
}
