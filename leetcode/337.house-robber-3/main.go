package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	cache := make(map[string]int)
	return singleRob(root, false, &cache)
}

func singleRob(node *TreeNode, skip bool, cache *map[string]int) int {
	if node == nil {
		return 0
	}

	chooseLeftKey := fmt.Sprintf("%p-choose", node.Left)
	chooseLeft, ok := (*cache)[chooseLeftKey]; if !ok {
		chooseLeft = singleRob(node.Left, false, cache)
		(*cache)[chooseLeftKey] = chooseLeft
	}

	chooseRightKey := fmt.Sprintf("%p-choose", node.Right)
	chooseRight, ok := (*cache)[chooseRightKey]; if !ok {
		chooseRight = singleRob(node.Right, false, cache)
		(*cache)[chooseRightKey] = chooseRight
	}
	skipThis := chooseLeft + chooseRight
	if skip {
		return skipThis
	}

	skipLeftKey := fmt.Sprintf("%p-skip", node.Left)
	skipLeft, ok := (*cache)[skipLeftKey]; if !ok {
		skipLeft = singleRob(node.Left, true, cache)
		(*cache)[skipLeftKey] = skipLeft
	}

	skipRightKey := fmt.Sprintf("%p-skip", node.Right)
	skipRight, ok := (*cache)[skipRightKey]; if !ok {
		skipRight = singleRob(node.Right, true, cache)
		(*cache)[skipRightKey] = skipRight
	}
	chooseThis := node.Val + skipLeft + skipRight

	return max(chooseThis, skipThis)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
