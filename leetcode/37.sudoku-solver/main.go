package main

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	board [][]byte
	rows  []uint16
	cols  []uint16
	cubes []uint16
	pos   [][2]int
)

func main() {
	board = [][]byte{
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
	solve(board)
	fmt.Printf("================================\n")
	printBlock()
	fmt.Println("done")
}

func solve(board [][]byte) error {
	initMask()
	dfs_v2(pos, rows, cols, cubes)
	return nil
}

func initMask() error {
	// init mark
	mc := make([]uint16, 9)
	mr := make([]uint16, 9)
	cs := make([]uint16, 9)
	ps := make([][2]int, 0)
	for i := range board {
		for j := range board[i] {
			if board[i][j] == '.' {
				ps = append(ps, [2]int{i, j})
				continue
			}
			num := board[i][j] - '0'
			mc[j] |= (1 << (num - 1))
			mr[i] |= (1 << (num - 1))
			cs[getCubeIndex(i, j)] |= (1 << (num - 1))
		}
	}
	rows = mr
	cols = mc
	cubes = cs
	pos = ps
	//printMask()
	return nil
}

func getCubeIndex(i, j int) int {
	return i/3*3 + j/3
}

func dfs_v2(pos [][2]int, rows, cols, cubes []uint16) bool {
	filled := make(map[int]struct{}, len(pos))
	counter := 0
	for len(filled) < len(pos) {
		prev := len(filled)
		//fmt.Printf("loop %d, pos len:%d, filled len:%d\n", counter, len(pos), len(filled))
		for p := range pos {
			if _, ok := filled[p]; ok {
				continue
			}
			i := pos[p][0]
			j := pos[p][1]
			c := getCubeIndex(i, j)
			pool := check(rows[i], cols[j], cubes[c])
			//fmt.Printf("pos:(%d,%d), cube:%d, available:%v\n", i, j, c, pool)
			if len(pool) == 0 {
				return false
			}
			if len(pool) == 1 {
				num := pool[0]
				bitMask := uint16(1 << (num - '1'))
				board[i][j] = num
				rows[i] |= bitMask
				cols[j] |= bitMask
				cubes[c] |= bitMask
				filled[p] = struct{}{}
				//fmt.Printf("fill %v to (%d,%d)\n", num-'0', i, j)
			}
		}
		if len(filled) == prev {
			// no new number
			//fmt.Printf("loop %d, pos len:%d, filled len:%d, no new number\n", counter, len(pos), len(filled))
			unsolved := make([][2]int, 0, len(pos)-len(filled))
			for idx := range pos {
				if _, ok := filled[idx]; ok {
					continue
				}
				unsolved = append(unsolved, pos[idx])
			}
			return dfs(unsolved, rows, cols, cubes)
		}
		counter++
	}
	return true
}

func dfs(pos [][2]int, rows, cols, cubes []uint16) bool {
	for p := range pos {
		i := pos[p][0]
		j := pos[p][1]
		c := getCubeIndex(i, j)
		pool := check(rows[i], cols[j], cubes[c])
		//fmt.Printf("row:%d, col:%d, cube:%d, available:%v\n", i, j, c, pool)
		if len(pool) == 0 {
			return false
		}
		for idx := range pool {
			num := pool[idx]
			bitMask := uint16(1 << (num - '1'))
			board[i][j] = num
			rows[i] |= bitMask
			cols[j] |= bitMask
			cubes[c] |= bitMask
			//fmt.Printf("fitting %d to (%d, %d)\n", num, i, j)
			//printBlock()
			if dfs(pos[p+1:], rows, cols, cubes) {
				return true
			}
			board[i][j] = '.'
			rows[i] ^= bitMask
			cols[j] ^= bitMask
			cubes[c] ^= bitMask
			//fmt.Printf("walk back to (%d, %d), available:%v\n", i, j, pool)
		}
		//fmt.Printf("no more choice\n")
		return false
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
	for i := range cubes {
		fmt.Printf("cubes %d, %v\n", i, uintToBits(cubes[i]))
	}
	fmt.Printf("-------------------\n")
	for i := range pos {
		fmt.Printf("pos %d, (%d, %d)", pos[i][0], pos[i][1])
	}
	fmt.Printf("-------------------\n")
}

func printBlock() {
	for i := range board {
		for j := range board[i] {
			if i%3 == 0 && j == 0 {
				fmt.Printf("-------------------\n")
			}
			var v interface{} = board[i][j] - '0'
			if board[i][j] == '.' {
				v = "."
			}
			if j%3 == 0 {
				fmt.Printf("|%v", v)
			} else {
				fmt.Printf(" %v", v)
			}
		}
		fmt.Printf("|\n")
	}
	fmt.Printf("-------------------\n")
}
