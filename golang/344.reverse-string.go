/*
 * @lc app=leetcode id=344 lang=golang
 *
 * [344] Reverse String
 */

// @lc code=start
func reverseString(s []byte)  {
    l := len(s)
    if l <= 1 {
		return
    }
    for i, j := 0, l - 1; i < j; i, j = i + 1, j - 1 {
		s[i], s[j] = s[j], s[i]
    }
}
// @lc code=end

