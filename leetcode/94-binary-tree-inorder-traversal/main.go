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

func(s *SimpleStack) Empty() bool {
	return len(s.Nodes) == 0
}

func inorderTraversal(root *TreeNode) []int {
	vals := make([]int, 0)
	stack := NewStack()
	cur := root

	for cur != nil || !stack.Empty() {
		for cur != nil {
			stack.Push(cur)
			cur = cur.Left
		}
		cur = stack.Pop()
		vals = append(vals, cur.Val)
		cur = cur.Right
	}

	return vals
}
