/*
 * @lc app=leetcode id=367 lang=golang
 *
 * [367] Valid Perfect Square
 */

// @lc code=start
func isPerfectSquare(num int) bool {
	if num < 0 {
		return false
	}
	left, right := num, 0
	for left > right {
		c := (left + right) / 2
		if c * c > num {
			left = c - 1
		} else if c * c < num {
			right = c + 1
		} else {
			return true
		}
	}
	
	return (left * left) == num
}
// @lc code=end

