package main

/**
 * s: input string, could be empty and contains only lowercase letters a-z.
 * p: pattern, could be empty and contains only lowercase letters a-z, and characters like . or *.
 *
 * '.' Matches any single character.
 * '*' Matches zero or more of the preceding element.
 * The matching should cover the entire input string (not partial).
 */
func isMatch(s string, p string) bool {
	memo := make(map[string]bool)
	return dp(s, p, 0, 0, memo)
}

func dp(s, p string, i, j int, memo map[string]bool) bool {
	if j >= len(p) {
		return i >= len(s)
	}
	key := fmt.Sprintf("%d-%d", i, j)
	if m, ok := memo[key]; ok {
		return m
	}
	match := i < len(s) && (p[j] == s[i] || p[j] == '.')
	var ret bool
	if j+1 < len(p) && p[j+1] == '*' {
		ret = (match && dp(s, p, i+1, j, memo)) || dp(s, p, i, j+2, memo)
	} else {
		ret =  match && dp(s, p, i+1, j+1, memo)
	}
	memo[key] = ret
	return ret
}
