package main

func uniquePaths(m int, n int) int {
	pre := make([]int, n)
	for j := 0; j < n; j++ {
		pre[j] = 1
	}
	cur := make([]int, n)
	copy(cur, pre)
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			cur[j] = pre[j] + cur[j-1]
		}
		pre = cur
	}
	return cur[n-1]
}

