package main

/**
 * define of optimal substructure
 *          root
 *         /    \
 *        l      r
 *       / \    / \
 *      ll  lr rl  rr
 * max(root) = max(root+ll+lr+rl+rr, l+r, l+rl+rr, r+ll+lr)
 * =>  max(root) = max(root+max(ll)+max(lr)+max(rl)+max(rr), max(l)+max(r))
 */

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

func (s *SimpleStack) Push(node *TreeNode) {
	s.Nodes = append(s.Nodes, node)
}

func (s *SimpleStack) Pop() *TreeNode {
	n := len(s.Nodes)
	if n > 0 {
		node := s.Nodes[n-1]
		s.Nodes = s.Nodes[0 : n-1]
		return node
	}
	return nil
}

func (s *SimpleStack) Peek() *TreeNode {
	n := len(s.Nodes)
	if n > 0 {
		node := s.Nodes[n-1]
		return node
	}
	return nil
}

func (s *SimpleStack) Empty() bool {
	return len(s.Nodes) == 0
}

func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}
	dp := make(map[*TreeNode]int)
	stack := NewStack()
	stack.Push(root)

	var cur = root
	var pre *TreeNode
	for !stack.Empty() {
		// postorder travelsal
		for cur != nil {
			stack.Push(cur)
			cur = cur.Left
		}
		cur = stack.Peek()
		if cur.Right == nil || cur.Right == pre {
			stack.Pop()
			dp[cur] = max(robNode(cur, dp), dp[cur.Left] + dp[cur.Right])
			pre = cur
			cur = nil
		} else {
			cur = cur.Right
		}
	}
	return dp[root]
}

func robNode(node *TreeNode, dp map[*TreeNode]int) int {
	sum := node.Val
	if node.Left != nil {
		if node.Left.Left != nil {
			sum += dp[node.Left.Left]
		}
		if node.Left.Right != nil {
			sum += dp[node.Left.Right]
		}
	}
	if node.Right != nil {
		if node.Right.Left != nil {
			sum += dp[node.Right.Left]
		}
		if node.Right.Right != nil {
			sum += dp[node.Right.Right]
		}
	}
	return sum
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
