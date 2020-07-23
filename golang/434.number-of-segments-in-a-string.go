/*
 * @lc app=leetcode id=434 lang=golang
 *
 * [434] Number of Segments in a String
 */

// @lc code=start
func countSegments(s string) int {
	state := 0 // 0: finding none empty, 1: finding empty
	count := 0
	for i := 0; i < len(s); i++ {
		if state == 0 && s[i] != ' ' {
			count++
			state = 1
		}

		if state == 1 && s[i] == ' ' {
			state = 0
		}
	}
	return count
}
// @lc code=end

