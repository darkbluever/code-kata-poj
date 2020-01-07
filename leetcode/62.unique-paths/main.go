package main

func uniquePaths(m int, n int) int {
	if m < 1 || n < 1 {
		return 0
	}
	cache := make(map[string]int)
	return singlePath(m, n, &cache)
}

func singlePath(m,n int, cache *map[string]int) int {
	if m == 1 || n == 1 {
		return 1
	}
	
	var up, left int
	var ok bool
	if up, ok = (*cache)[fmt.Sprintf("%d-%d", m, n-1)]; !ok {
		up = singlePath(m, n-1, cache)
		(*cache)[fmt.Sprintf("%d-%d", m, n-1)] = up
	}
	if left, ok = (*cache)[fmt.Sprintf("%d-%d", m-1, n)]; !ok {
		left = singlePath(m-1, n, cache)
		(*cache)[fmt.Sprintf("%d-%d", m-1, n)] = left
	}
	return left + up
}
