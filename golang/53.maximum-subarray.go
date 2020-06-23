/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 */

// @lc code=start
func maxSubArray(nums []int) int {
	tl := len(nums)
	if tl == 0 {
		return 0
	}

	max, sum := nums[0], nums[0]
	for i := 1; i < tl; i++ {
		t := sum + nums[i]
		if t > nums[i] {
			sum += nums[i]
		} else {
			sum = nums[i]
		}

		if sum > max {
			max = sum
		}
	}

	return max
}
// @lc code=end

