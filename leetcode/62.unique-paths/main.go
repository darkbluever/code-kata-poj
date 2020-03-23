package main

func uniquePaths(m int, n int) int {
	cur := make([]int, n)
	for j := 0; j < n; j++ {
		cur[j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			cur[j] += cur[j-1]
		}
	}
	return cur[n-1]
}

