/*
 * @lc app=leetcode id=225 lang=golang
 *
 * [225] Implement Stack using Queues
 */

// @lc code=start
type MyStack struct {
	q1 []int
	q2 []int
	qc  *[]int
	qb 	*[]int
}


/** Initialize your data structure here. */
func Constructor() MyStack {
    s := MyStack {
		q1: 	make([]int, 0),
		q2: 	make([]int, 0),
		qc: 	nil,
		qb: 	nil,
	}
	s.qc = &s.q1
	s.qb = &s.q2
	return s
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
	*(this.qc) = append(*(this.qc), x)
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	l := len(*(this.qc))
	*(this.qb) = (*(this.qc))[:l - 1]
	ret := (*(this.qc))[l - 1]
	*(this.qc) = (*(this.qc))[:0]
	this.qc, this.qb = this.qb, this.qc
	return ret
}


/** Get the top element. */
func (this *MyStack) Top() int {
	l := len(*(this.qc))
    return (*(this.qc))[l - 1]
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
    return len(*(this.qc)) == 0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
// @lc code=end

