/*
 * @lc app=leetcode id=198 lang=golang
 *
 * [198] House Robber
 */

// @lc code=start
func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rob(nums []int) int {
	max_last, max := 0, 0
	for _, v := range nums {
		t := max
		max = maxInt(v + max_last, max)
		max_last = t
	}
	return max
}

// @lc code=end

