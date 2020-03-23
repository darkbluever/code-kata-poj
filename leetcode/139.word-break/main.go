package main

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s) + 1)
	dp[0] = true
	wordMap := make(map[string]struct{}, len(wordDict))
	for k := range wordDict {
		wordMap[wordDict[k]] = struct{}{}
	}

	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			if _, ok := wordMap[s[i:j]]; dp[i] && ok {
				dp[j] = true
			}
		}
	}

	return dp[len(s)]
}
