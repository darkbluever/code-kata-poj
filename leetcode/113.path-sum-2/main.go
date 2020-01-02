package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return [][]int{}
	}
	if root.Left == nil && root.Right == nil {
		if root.Val == sum {
			return [][]int{[]int{root.Val}}
		}
		return [][]int{}
	}
	ret := make([][]int, 0)
	left := pathSum(root.Left, sum-root.Val)
	right := pathSum(root.Right, sum-root.Val)
	if len(left) > 0 {
		for i := range left {
			tmp := make([]int, 0, len(left[i])+1)
			tmp = append(tmp, root.Val)
			tmp = append(tmp, left[i]...)
			left[i] = tmp
		}
		ret = append(ret, left...)
	}
	if len(right) > 0 {
		for i := range right {
			tmp := make([]int, 0, len(right[i])+1)
			tmp = append(tmp, root.Val)
			tmp = append(tmp, right[i]...)
			right[i] = tmp
		}
		ret = append(ret, right...)
	}

	return ret
}
