/*
 * @lc app=leetcode id=69 lang=golang
 *
 * [69] Sqrt(x)
 */

// @lc code=start
func mySqrt(x int) int {
    if x < 0 {
		return -1
	} else if x == 1 || x == 0 {
		return x
	}

	t := int64(x)
	left, right := t, int64(0)
	for left > right + 1 {
		c := (left + right) / 2
		if c * c == t {
			return int(c)
		} else if c * c > t {
			left = c
		} else {
			right = c
		}
	}

	return int(right)
}
// @lc code=end

