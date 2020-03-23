package main

func climbStairs(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 {
		return 1
	}
	first, second := 1, 2
	for i := 3; i <= n; i++ {
		second = first + second
		first = second - first
	}
	return second
}

