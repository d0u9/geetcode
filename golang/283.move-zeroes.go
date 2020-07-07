/*
 * @lc app=leetcode id=283 lang=golang
 *
 * [283] Move Zeroes
 */

// @lc code=starts
// Solved via internet searching
func moveZeroes(nums []int)  {
    for i, j := 0, 0; i < len(nums); i++ {
        if nums[i] != 0 {
            nums[j], nums[i] = nums[i], nums[j]
            j++
        }
    }
}
// @lc code=end

