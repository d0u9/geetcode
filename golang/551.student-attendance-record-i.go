/*
 * @lc app=leetcode id=551 lang=golang
 *
 * [551] Student Attendance Record I
 */

// @lc code=start
func checkRecord(s string) bool {
	l := len(s)
	ca, cl := 0, 0
	for i := 0; i < l; i++ {
		if s[i] == 'A' {
			ca++
			cl = 0
			if ca > 1 {
				return false
			}
		} else if s[i] == 'L' {
			cl++
			if cl > 2 {
				return false
			}
		} else {
			cl = 0
		}
	}
    return true
}
// @lc code=end

