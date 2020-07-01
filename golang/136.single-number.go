/*
 * @lc app=leetcode id=136 lang=golang
 *
 * [136] Single Number
 */

// @lc code=start
func singleNumber(nums []int) int {
	ret := 0
	for _, v := range nums {
		ret = ret ^ v
	}
    return ret
}
// @lc code=end

