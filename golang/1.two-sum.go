/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	l := len(nums)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if nums[i] + nums[j] == target {
				return []int{i, j}
			}
		}
	}
  
	return []int{0, 0}
}
// @lc code=end

