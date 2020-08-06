/*
 * @lc app=leetcode id=561 lang=golang
 *
 * [561] Array Partition I
 */

// @lc code=start
func arrayPairSum(nums []int) int {
	l := len(nums)
	sort.Ints(nums)
	sum := 0
	for i := 0; i < l; i = i + 2 {
		sum = sum + nums[i]
	}
    return sum
}
// @lc code=end

