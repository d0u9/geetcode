/*
 * @lc app=leetcode id=232 lang=golang
 *
 * [232] Implement Queue using Stacks
 */

// @lc code=start
type MyQueue struct {
	sc 	*[]int
	sb 	*[]int
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
	s1, s2 := make([]int, 0), make([]int, 0)
    return MyQueue {
		sc: &s1,
		sb: &s2,
	}
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
    *(this.sc) = append(*(this.sc), x)
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	l := len(*(this.sc))
	for i := l - 1; i > 0; i-- {
		v := (*(this.sc))[i]
		*(this.sb) = append(*(this.sb), v)
	}
	ret := (*(this.sc))[0]
	*(this.sc) = (*(this.sc))[:0]

	l = len(*(this.sb))
	for i := l - 1; i >= 0; i-- {
		v := (*(this.sb))[i]
		*(this.sc) = append(*(this.sc), v)
	}
	*(this.sb) = (*(this.sb))[:0]
	fmt.Printf("%v, %v\n", *(this.sc), *(this.sb))
	return ret
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
    return (*(this.sc))[0]
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(*(this.sc)) == 0
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
// @lc code=end

