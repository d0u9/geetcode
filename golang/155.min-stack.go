/*
 * @lc app=leetcode id=155 lang=golang
 *
 * [155] Min Stack
 */

// @lc code=start
type Pair struct {
	val int
	min int
}

type MinStack struct {
	vals []Pair
	min int
	cur int
}


/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack {
		vals: make([]Pair, 0),
		min: math.MaxInt32,
		cur: 0,
	}
}

func (this *MinStack) Push(x int)  {
	if x < this.min {
		this.min = x
	}
	this.vals = append(this.vals, Pair{ val: x, min: this.min, })
	this.cur++
}


func (this *MinStack) Pop()  {
	this.cur--
	this.vals = this.vals[:this.cur]
	//fmt.Println(this.vals)
	if this.cur == 0 {
		this.min = math.MaxInt32
	} else {
		this.min = this.vals[this.cur - 1].min
	}
}


func (this *MinStack) Top() int {
    return this.vals[this.cur - 1].val
}


func (this *MinStack) GetMin() int {
    return this.min
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end

