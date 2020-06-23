/*
 * @lc app=leetcode id=58 lang=golang
 *
 * [58] Length of Last Word
 */

// @lc code=start
func lengthOfLastWord(s string) int {
	tl := len(s)
	if tl == 0 {
		return 0
	}
	i, j := 0, 0
	for i = tl - 1; i >= 0; i-- {
		if s[i] != ' ' {
			break
		}
	}

	for j = i; j >= 0; j-- {
		if s[j] == ' ' {
			break
		}
	}

	return i - j
}
// @lc code=end

