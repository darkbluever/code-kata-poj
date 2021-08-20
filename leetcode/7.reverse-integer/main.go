package main

import (
	"fmt"
	"math"
)

func main() {
	num := 123
	fmt.Printf("input:%d, output:%d\n", num, reverse(num))
}

func reverse(x int) int {
	ret := 0
	for x > 0 {
		if ret > math.MaxInt32 / 10 || ret < math.MinInt32 / 10 {
			return 0
		}
		d := x % 10
		x = x / 10
		ret = ret*10 + d
	}
	return ret
}
