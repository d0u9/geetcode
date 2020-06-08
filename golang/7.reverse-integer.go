/*
 * @lc app=leetcode id=7 lang=golang
 *
 * [7] Reverse Integer
 */

// @lc code=start
func reverse(x int) int {
	y := int((1 << 31) - 1)
	z := int(0 - (1 << 31)) 

	negative := 1
	value := x
	if x < 0 {
		negative = -1
		value = -x
	}

	output := 0

	for value != 0 {
		rem := value % 10
		output = output * 10 + rem
		value = value / 10
	}

	k := output * negative
	if k > y || k < z {
		return 0
	}
	return k
}
// @lc code=end

