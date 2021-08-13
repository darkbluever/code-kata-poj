package main

import "sort"

func hIndex(citations []int) int {
	l := len(citations)
	sort.Ints(citations)
	for i := range citations {
		if citations[i] >= l-i {
			return l - i
		}
	}
	return 0
}
