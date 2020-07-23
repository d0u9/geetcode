/*
 * @lc app=leetcode id=485 lang=golang
 *
 * [485] Max Consecutive Ones
 */

// @lc code=start
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findMaxConsecutiveOnes(nums []int) int {
	l := len(nums)
	if l <= 0 {
		return 0
	}

	cnt, max := 0, 0
	for i := 0; i < l; i++ {
		if nums[i] != 1 {
			max = Max(max, cnt)
			cnt = 0
		} else {
			cnt++
		}
	}

	max = Max(max, cnt)

	return max
}
// @lc code=end

