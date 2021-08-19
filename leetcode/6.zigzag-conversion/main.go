package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "PAYPALISHIRING"
	num := 4

	fmt.Printf("input: %s, num: %d\n", str, num)
	ret := convert(str, num)
	fmt.Printf("output: %s\n", ret)
}

func convert(s string, numRows int) string {
	if numRows <= 1 || len(s) <= numRows {
		return s
	}
	var parts = make([]string, numRows)
	for i := range s {
		pos := i % (numRows + numRows - 2)
		if pos < numRows {
			parts[pos] += string(s[i])
		} else {
			parts[2*(numRows-1)-pos] += string(s[i])
		}
	}
	return strings.Join(parts, "")
}
