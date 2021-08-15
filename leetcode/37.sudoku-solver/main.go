package main

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	block [][]byte
	rows []uint16
	cols []uint16
	cubes []uint16
)

func main() {
	block = [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	fmt.Printf("input:\n")
	printBlock()
	solve(block)
	fmt.Printf("================================\n")
	printBlock()
	fmt.Println("done")
}

func solve(block [][]byte) error {
	initMask()
	dfs(block, rows, cols, cubes, 0, 0)
	return nil
}

func initMask() error {
	// init mark
	mc := make([]uint16, 9)
	mr := make([]uint16, 9)
	cs := make([]uint16, 9)
	for i := range block {
		for j := range block[i] {
			if block[i][j] == '.' {
				continue
			}
			num := block[i][j] - '0'
			mc[j] |= (1 << (num - 1))
			mr[i] |= (1 << (num - 1))
			cs[getCubeIndex(i, j)] |= (1 << (num - 1))
		}
	}
	rows = mr
	cols = mc
	cubes = cs
	//printMask()
	return nil
}

func getCubeIndex(i, j int) int {
	h := i / 3
	v := j / 3
	return h*3 + v
}

func dfs(block [][]byte, rows, cols, cubes []uint16, x, y int) bool {
	if x > len(rows) || y > len(cols) {
		//fmt.Printf("invalid params, x:%d, y:%d\n", x, y)
		return false
	}
	for i := x; i < len(block); i++ {
		if i > x {
			y = 0
		}
		for j := y; j < len(block[i]); j++ {
			if block[i][j] == '.' {
				c := getCubeIndex(i, j)
				pool := check(rows[i], cols[j], cubes[c])
				//fmt.Printf("row:%d, col:%d, cube:%d, available:%v\n", i, j, c, pool)
				if len(pool) == 0 {
					return false
				}
				for idx := range pool {
					num := pool[idx]
					bitMask := uint16(1 << (num - '1'))
					block[i][j] = num
					rows[i] |= bitMask
					cols[j] |= bitMask
					cubes[c] |= bitMask
					//fmt.Printf("fitting %d to (%d, %d)\n", num, i, j)
					//printBlock()
					m, n := i, j+1
					if n == len(cols) {
						n = 0
						m = i + 1
					}
					if dfs(block, rows, cols, cubes, m, n) {
						return true
					}
					block[i][j] = '.'
					rows[i] ^= bitMask
					cols[j] ^= bitMask
					cubes[c] ^= bitMask
					//debugMask(i, j)
					//fmt.Printf("walk back to (%d, %d), available:%v\n", i, j, pool)
				}
				//fmt.Printf("no more choice\n")
				return false
			}
		}
	}
	return true
}

func check(row, col, cube uint16) []byte {
	available := make([]byte, 0, 9)
	ret := row | col | cube
	for i := byte(0); i < 9; i++ {
		n := (ret >> i) & 1
		if n == 0 {
			available = append(available, i+'1')
		}
	}
	return available
}

func uintToBits(num uint16) string {
	var sb strings.Builder
	for i := 0; i < 9; i++ {
		move := uint(8 - i)
		sb.WriteString(strconv.FormatInt(int64((num>>move)&1), 10))
	}

	return sb.String()
}

func printMask() {
	fmt.Printf("-------------------\n")
	for i := range rows {
		fmt.Printf("row %d, %v\n", i, uintToBits(rows[i]))
	}
	fmt.Printf("-------------------\n")
	for i := range cols {
		fmt.Printf("col %d, %v\n", i, uintToBits(cols[i]))
	}
	fmt.Printf("-------------------\n")
	for i := range cols {
		fmt.Printf("cubes %d, %v\n", i, uintToBits(cubes[i]))
	}
	fmt.Printf("-------------------\n")
}

func printBlock() {
	fmt.Printf("-------------------\n")
	fmt.Printf("|%v %v %v|%v %v %v|%v %v %v|\n", string(block[0][0]), string(block[0][1]), string(block[0][2]), string(block[0][3]), string(block[0][4]), string(block[0][5]), string(block[0][6]), string(block[0][7]), string(block[0][8]))
	fmt.Printf("|%v %v %v|%v %v %v|%v %v %v|\n", string(block[1][0]), string(block[1][1]), string(block[1][2]), string(block[1][3]), string(block[1][4]), string(block[1][5]), string(block[1][6]), string(block[1][7]), string(block[1][8]))
	fmt.Printf("|%v %v %v|%v %v %v|%v %v %v|\n", string(block[2][0]), string(block[2][1]), string(block[2][2]), string(block[2][3]), string(block[2][4]), string(block[2][5]), string(block[2][6]), string(block[2][7]), string(block[2][8]))
	fmt.Printf("-------------------\n")
	fmt.Printf("|%v %v %v|%v %v %v|%v %v %v|\n", string(block[3][0]), string(block[3][1]), string(block[3][2]), string(block[3][3]), string(block[3][4]), string(block[3][5]), string(block[3][6]), string(block[3][7]), string(block[3][8]))
	fmt.Printf("|%v %v %v|%v %v %v|%v %v %v|\n", string(block[4][0]), string(block[4][1]), string(block[4][2]), string(block[4][3]), string(block[4][4]), string(block[4][5]), string(block[4][6]), string(block[4][7]), string(block[4][8]))
	fmt.Printf("|%v %v %v|%v %v %v|%v %v %v|\n", string(block[5][0]), string(block[5][1]), string(block[5][2]), string(block[5][3]), string(block[5][4]), string(block[5][5]), string(block[5][6]), string(block[5][7]), string(block[5][8]))
	fmt.Printf("-------------------\n")
	fmt.Printf("|%v %v %v|%v %v %v|%v %v %v|\n", string(block[6][0]), string(block[6][1]), string(block[6][2]), string(block[6][3]), string(block[6][4]), string(block[6][5]), string(block[6][6]), string(block[6][7]), string(block[6][8]))
	fmt.Printf("|%v %v %v|%v %v %v|%v %v %v|\n", string(block[7][0]), string(block[7][1]), string(block[7][2]), string(block[7][3]), string(block[7][4]), string(block[7][5]), string(block[7][6]), string(block[7][7]), string(block[7][8]))
	fmt.Printf("|%v %v %v|%v %v %v|%v %v %v|\n", string(block[8][0]), string(block[8][1]), string(block[8][2]), string(block[8][3]), string(block[8][4]), string(block[8][5]), string(block[8][6]), string(block[8][7]), string(block[8][8]))
	fmt.Printf("-------------------\n")
}
