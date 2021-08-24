package main

import (
	"fmt"
	"strings"
)

func main() {
	testCases := []struct{
		num int
		expect string
	} {
		{
			num: 3,
			expect: "III",
		},
		{
			num: 4,
			expect: "IV",
		},
		{
			num: 9,
			expect: "IX",
		},
		{
			num: 58,
			expect: "LVIII",
		},
		{
			num: 1994,
			expect: "MCMXCIV",
		},
	}
	for _, c := range testCases {
		fmt.Printf("int: %d, roman: %s, expect: %s\n", c.num, intToRoman(c.num), c.expect)
	}
}

func intToRoman(num int) string {
	nums := [13]int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	roman := [13]string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	var sb strings.Builder
	for i := range nums {
		for num >= nums[i] {
			sb.WriteString(roman[i])
			num -= nums[i]
		}
	}
	return sb.String()
}
