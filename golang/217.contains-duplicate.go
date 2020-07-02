/*
 * @lc app=leetcode id=217 lang=golang
 *
 * [217] Contains Duplicate
 */

// @lc code=start
func solution1(nums []int) bool {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i - 1] == nums[i] {
			return true
		}
	}
	return false
}

func solution2(nums []int) bool {
	mp := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if _, ok := mp[nums[i]]; !ok {
			mp[nums[i]] = true
		} else {
			return true
		}
	}
	return false
}

func containsDuplicate(nums []int) bool {
	return solution2(nums)
}
// @lc code=end

