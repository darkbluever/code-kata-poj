package main

func climbStairs(n int) int {
	cache := make(map[int]int)
	return climb(n, &cache)
}

func climb(n int, cache *map[int]int) int {
	if n < 1 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	var a, b int
	var ok bool
	if a, ok = (*cache)[n-1]; !ok {
		a = climb(n-1, cache)
		(*cache)[n-1] = a
	}
	if b, ok = (*cache)[n-2]; !ok {
		b = climb(n-2, cache)
		(*cache)[n-2] = b
	}
	return a+b
}
