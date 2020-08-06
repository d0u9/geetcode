/*
 * @lc app=leetcode id=521 lang=golang
 *
 * [521] Longest Uncommon Subsequence I 
 */

// @lc code=start

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findLUSlength(a string, b string) int {
	la, lb := len(a), len(b)

	if a == b {
		return -1
	}

	return MaxInt(la, lb)
}
// @lc code=end

