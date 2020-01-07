package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	ret := make([][]int, 0)
	path := make([]int, 0)
	singlePathSum(root, sum, path, &ret)

	return ret
}

func singlePathSum(node *TreeNode, sum int, path []int, result *[][]int) {
	if node == nil {
		return
	}
	path = append(path, node.Val)
	if node.Left == nil && node.Right == nil && node.Val == sum {
		ret := make([]int, len(path))
		copy(ret, path)
		*result = append(*result, ret)
		return
	}
	singlePathSum(node.Left, sum-node.Val, path, result)
	singlePathSum(node.Right, sum-node.Val, path, result)
}
