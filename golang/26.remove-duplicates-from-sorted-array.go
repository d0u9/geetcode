/*
 * @lc app=leetcode id=26 lang=golang
 *
 * [26] Remove Duplicates from Sorted Array
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}

	ret := 1
	for i, j := 0, 1; i < l && j < l; {
		if nums[j] == nums[i] {
			j++
		} else {
			i++
			nums[i] = nums[j]
			j++
			ret++	
		}
	}

	return ret
}
// @lc code=end

