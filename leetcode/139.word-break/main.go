package main

var a = rune('a')

type TreeNode struct {
	count    int
	children [26]*TreeNode
}

func NewTireTree() *TreeNode {
	return &TreeNode{}
}

func (t *TreeNode) Add(s string) {
	if len(s) == 0 {
		return
	}
	r := []rune(s)
	cur := t
	for i := range r {
		idx := r[i] - a
		if cur.children[idx] == nil {
			cur.children[idx] = NewTireTree()
		}
		cur = cur.children[idx]
	}
	cur.count = 1
}

func (t *TreeNode) Get(s string) int {
	if len(s) == 0 {
		return 1
	}
	r := []rune(s)
	cur := t
	for i := range r {
		idx := r[i] - a
		if cur.children[idx] == nil {
			return -1
		}
		cur = cur.children[idx]
	}
	return cur.count
}

func wordBreak(s string, wordDict []string) bool {
	if len(wordDict) == 0 || len(s) == 0 {
		return false
	}
	root := NewTireTree()
	for i := range wordDict {
		root.Add(wordDict[i])
	}
	return walk(s, root)
}

func walk(s string, t *TreeNode) bool {
	flag := false
	for i := range s {
		code := t.Get(s[0:i+1])
		if code == -1 {
			return flag
		}
		if code == 1 {
			if i == len(s) - 1 {
				return true
			}
			b := walk(s[i+1:], t)
			flag = flag || b
		}
	}
	return flag
}
