package main

func findJudge(N int, trust [][]int) int {
	in := make([]int, N+1)
	out := make([]int, N+1)
	for i := range trust {
		out[trust[i][0]]++
		in[trust[i][1]]++
	}

	for i := 1; i <= N; i++ {
		if out[i] == 0 && in[i] == N-1 {
			return i
		}
	}
	return -1
}
