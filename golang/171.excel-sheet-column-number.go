/*
 * @lc app=leetcode id=171 lang=golang
 *
 * [171] Excel Sheet Column Number
 */

// @lc code=start
func titleToNumber(s string) int {
	ret := 0
	base := 1
	for i := len(s) - 1; i >= 0; i-- {
		ret = ret + ((int(s[i] - 'A') + 1)) * base
		base = base * 26
	}
    return ret
}
// @lc code=end

