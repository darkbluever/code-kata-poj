package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type SimpleStack struct {
	Nodes []*TreeNode
}

func NewStack() *SimpleStack {
	return &SimpleStack{[]*TreeNode{}}
}

func(s *SimpleStack) Push(node *TreeNode) {
	s.Nodes = append(s.Nodes, node)
}

func(s *SimpleStack) Pop() *TreeNode {
	n := len(s.Nodes)
	if n > 0 {
		node := s.Nodes[n-1]
		s.Nodes = s.Nodes[0:n-1]
		return node
	}
	return nil
}

func inorderTraversal(root *TreeNode) []int {
	vals := make([]int, 0)
	if root == nil {
		return vals
	}
	stack := NewStack()
	cur := root

	for cur != nil {
		if cur.Left != nil {
			stack.Push(cur)
			cur = cur.Left
			continue
		}
		vals = append(vals, cur.Val)
		if cur.Right != nil {
			cur = cur.Right
			continue
		}
		cur = stack.Pop()
		for cur != nil {
			vals = append(vals, cur.Val)
			if cur.Right != nil {
				cur = cur.Right
				break
			}
			cur = stack.Pop()
		}
	}

	return vals
}
