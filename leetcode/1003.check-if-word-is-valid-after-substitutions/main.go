package main

const (
	ByteA = byte('a')
	ByteB = byte('b')
	ByteC = byte('c')
)

type stack struct {
	buf []byte
}

func NewStack() *stack {
	return &stack{make([]byte, 0, 0)}
}

func (s *stack) Push(b byte) {
	s.buf = append(s.buf, b)
}

func (s *stack) Pop() byte {
	length := len(s.buf)
	if length == 0 {
		return 0
	}
	b := s.buf[length-1]
	s.buf = s.buf[:length-1]
	return b
}

func (s *stack) Empty() bool {
	return len(s.buf) == 0
}

func isValid(S string) bool {
	ret := true
	stack := NewStack()
	for i := range S {
		if S[i] == ByteC {
			if stack.Pop() == ByteB && stack.Pop() == ByteA {
				continue
			}
			ret = false
			break
		}
		stack.Push(S[i])
	}
	if !stack.Empty() {
		ret = false
	}
	return ret
}
