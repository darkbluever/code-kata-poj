package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	return singleRob(root, true)
}

func singleRob(node *TreeNode, flag bool) int {
	if node == nil {
		return 0
	}
	if flag {
		choose := node.Val + singleRob(node.Left, false) + singleRob(node.Right, false)
		skip := singleRob(node.Left, true) + singleRob(node.Right, true)
		return max(choose, skip)
	}
	return singleRob(node.Left, true) + singleRob(node.Right, true)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
