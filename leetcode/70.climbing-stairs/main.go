package main

func climbStairs(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	first, second := 1, 2
	sum = first + second
	for i := 3; i < n; i++ {
		sum = sum + second
		first, second = second, sum - second
	}
	return sum
}

