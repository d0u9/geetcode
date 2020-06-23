/*
 * @lc app=leetcode id=35 lang=golang
 *
 * [35] Search Insert Position
 */

// @lc code=start
import "fmt"
func searchInsert(nums []int, target int) int {
	tl := len(nums)
	l, r := 0, tl - 1

	for l <= r {
		c := int((l + r) / 2)
		if nums[c] < target {
			l = c + 1
		} else if nums[c] > target {
			r = c - 1
		} else {
			return c
		}
	}

	return l
}
// @lc code=end

