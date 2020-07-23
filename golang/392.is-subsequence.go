/*
 * @lc app=leetcode id=392 lang=golang
 *
 * [392] Is Subsequence
 */

// @lc code=start
func isSubsequence(s string, t string) bool {
	l1, l2 := len(s), len(t)
	for i, j := 0, 0; i < l1; i++ {
		if j >= l2 {
			return false
		}

		c := s[i]
		find := false

		for ;j < l2;{
			v := t[j]
			j++
			if v == c {
				find = true
				break
			}
		}

		if !find {
			return false
		}
	}
	return true
}
// @lc code=end

