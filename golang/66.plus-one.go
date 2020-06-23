/*
 * @lc app=leetcode id=66 lang=golang
 *
 * [66] Plus One
 */

// @lc code=start
func plusOne(digits []int) []int {
	tl := len(digits)
	if tl == 0 {
		return digits
	}

	i, c := 0, 1
	for i = tl - 1; i >= 0; i-- {
		r := digits[i] + c
		d := r % 10
		c  = r / 10
		digits[i] = d
	}

	if c == 0 {
		return digits
	}

	r := []int {c, }
	return append(r, digits...)
}
// @lc code=end

