package main

import (
	"fmt"
	"math"
)

func main() {
	testCases := []struct {
		input  string
		expect int
	}{
		{"42", 42},
		{"    -42", -42},
		{"4193 with words", 4193},
		{"words and 987", 0},
		{"-91283472332", -2147483648},
	}

	for i, c := range testCases {
		ret := myAtoiDFA(c.input)
		fmt.Printf("test case [%d], input:%s, output:%d, expect:%d\n", i, c.input, ret, c.expect)
	}
}

//函数 myAtoi(string s) 的算法如下：
//读入字符串并丢弃无用的前导空格
//检查下一个字符（假设还未到字符末尾）为正还是负号，读取该字符（如果有）。 确定最终结果是负数还是正数。 如果两者都不存在，则假定结果为正。
//读入下一个字符，直到到达下一个非数字字符或到达输入的结尾。字符串的其余部分将被忽略。
//将前面步骤读入的这些数字转换为整数（即，"123" -> 123， "0032" -> 32）。如果没有读入数字，则整数为 0 。必要时更改符号（从步骤 2 开始）。
//如果整数数超过 32 位有符号整数范围 [−2^31,  2^31 − 1] ，需要截断这个整数，使其保持在这个范围内。具体来说，小于 −2^31 的整数应该被固定为 −2^31 ，大于 2^31 − 1 的整数应该被固定为 2^31 − 1 。
//返回整数作为最终结果。

//注意：
//本题中的空白字符只包括空格字符 ' ' 。
//除前导空格或数字后的其余字符串外，请勿忽略 任何其他字符。
func myAtoi(s string) int {
	leadingSpace := true
	sign := true
	base := 1
	ret := 0
	for i := range s {
		if leadingSpace {
			if s[i] == ' ' {
				continue
			} else {
				leadingSpace = false
			}
		}
		if sign {
			if s[i] == '-' {
				sign = false
				base = -1
				continue
			} else if s[i] == '+' {
				sign = false
				continue
			}
			sign = false
		}
		if s[i] < '0' || s[i] > '9' {
			break
		}
		ret = ret*10 + int(s[i]-'0')*base
		if ret > math.MaxInt32 {
			return math.MaxInt32
		}
		if ret < math.MinInt32 {
			return math.MinInt32
		}
	}
	return ret
}

func myAtoiDFA(s string) int {
	dfa := newDFA()
	for i := range s {
		dfa.read(s[i])
	}
	return dfa.v
}

type status uint8

const (
	statusStart status = iota
	statusSigned
	statusNumber
	statusEnd
)

type dfa struct {
	s    status
	m    map[status][]status
	sign int
	v    int
}

func newDFA() *dfa {
	return &dfa{
		s: statusStart,
		m: map[status][]status{
			statusStart:  []status{statusStart, statusSigned, statusNumber, statusEnd},
			statusSigned: []status{statusEnd, statusEnd, statusNumber, statusEnd},
			statusNumber: []status{statusEnd, statusEnd, statusNumber, statusEnd},
			statusEnd:    []status{statusEnd, statusEnd, statusEnd, statusEnd},
		},
		sign: 1,
	}
}

func (a *dfa) read(char byte) {
	s := a.m[a.s][a.nextStatus(char)]
	a.s = s
	switch s {
	case statusSigned:
		if char == '-' {
			a.sign = -1
		}
	case statusNumber:
		a.v = a.v*10 + int(char-'0')*a.sign
		if a.v > math.MaxInt32 {
			a.s = statusEnd
			a.v = math.MaxInt32
		}
		if a.v < math.MinInt32 {
			a.s = statusEnd
			a.v = math.MinInt32
		}
	}
}

func (a *dfa) nextStatus(char byte) int {
	if char == ' ' {
		return 0
	}
	if char == '-' || char == '+' {
		return 1
	}
	if char >= '0' && char <= '9' {
		return 2
	}
	return 3
}
