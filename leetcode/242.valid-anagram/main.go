package main

func isAnagram(s string, t string) bool {
	counter := make([]int, 26)
	for _, v := range s {
		idx := v - 'a'
		counter[idx] += 1
	}
	for _, v := range t {
		idx := v - 'a'
		counter[idx] -= 1
	}
	for i := range counter {
		if counter[i] != 0 {
			return false
		}
	}
	return true
}
