package main

type MinStack struct {
	stack   []int
	min     int64
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack:   make([]int, 0, 0),
		min:     math.MaxInt64,
	}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	if this.min > int64(x) {
		this.min = int64(x)
	}
}

func (this *MinStack) Pop() {
	tmp := this.stack[len(this.stack) - 1]
	this.stack = this.stack[0 : len(this.stack)-1]
	if this.min == int64(tmp) {
		this.min = math.MaxInt64
		for i := range this.stack {
			if this.min > int64(this.stack[i]) {
				this.min = int64(this.stack[i])
			}
		}
	}
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return int(this.min)
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
