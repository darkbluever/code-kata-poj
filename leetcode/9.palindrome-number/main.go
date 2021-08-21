package main

import "fmt"

func main() {
	testCases := []int{121, -121, 12321, 1234, 1221, 0, 10, 1000030001}
	for _, c := range testCases {
		fmt.Printf("%d is palindrome? %v\n", c, isPalindrome(c))
	}
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	var digit = make([]int, 0)
	for x != 0 {
		digit = append(digit, x%10)
		x = x / 10
	}
	start, end := 0, len(digit)-1
	for start < end {
		if digit[start] != digit[end] {
			return false
		}
		start++
		end--
	}
	return true
}
