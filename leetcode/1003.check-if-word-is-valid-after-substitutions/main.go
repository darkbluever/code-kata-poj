package main

func isValid(S string) bool {
	for {
		if strings.Contains(S, "abc") {
			S = strings.ReplaceAll(S, "abc", "")
			continue
		}
		break
	}
	return len(S) == 0
}
