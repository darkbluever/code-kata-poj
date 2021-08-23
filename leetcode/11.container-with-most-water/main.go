package main

import (
	"fmt"
)

func main() {
	testCases := []struct{
		height []int
		expect int
	}{
		{
			height: []int{1,8,6,2,5,4,8,3,7},
			expect: 49,
		},
		{
			height: []int{1,1},
			expect: 1,
		},
		{
			height: []int{4,3,2,1,4},
			expect: 16,
		},
		{
			height: []int{1,2,1},
			expect: 2,
		},
	}
	for _, c := range testCases {
		fmt.Printf("input:%v, max area:%d, expect:%d\n", c.height, maxArea(c.height), c.expect)
	}
}


func maxArea(height []int) int {
	start, end := 0, len(height) -1
	area := 0
	for start < end {
		tmp := (end - start) * min(height[start], height[end])
		if tmp > area {
			area = tmp
		}
		if height[start] < height[end] {
			start++
		} else {
			end--
		}
	}
	return area
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
