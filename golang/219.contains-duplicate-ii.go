/*
 * @lc app=leetcode id=219 lang=golang
 *
 * [219] Contains Duplicate II
 */

// @lc code=start
func solution1(nums []int, k int) bool {
	mp := make(map[int]int)
	for i, v := range nums {
		if j, ok := mp[v]; ok {
			if i - j <= k {
				return true
			}
		}
		mp[v] = i
	}
	return false
}



func containsNearbyDuplicate(nums []int, k int) bool {
    return solution1(nums, k)
}
// @lc code=end

