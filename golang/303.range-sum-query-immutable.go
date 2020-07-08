/*
 * @lc app=leetcode id=303 lang=golang
 *
 * [303] Range Sum Query - Immutable
 */

// @lc code=start
type NumArray struct {
    s []int
}


func Constructor(nums []int) NumArray {
	l, sum := len(nums), 0
	s := make([]int, l + 1, l + 1)
	s[0] = 0
    for i := 0; i < l; i++ {
		sum += nums[i]
		s[i + 1] = sum
	}
	return NumArray {
		s: s,
	}
}


func (this *NumArray) SumRange(i int, j int) int {
    return this.s[j + 1] - this.s[i]
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
// @lc code=end

