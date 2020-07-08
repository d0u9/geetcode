/*
 * @lc app=leetcode id=342 lang=golang
 *
 * [342] Power of Four
 */

// @lc code=start
func solution1(num int) bool {
	if num < 0 {
		return false
	} else if num == 1 {
		return true
	}
	for i := 1; i < 16; i++ {
		if num == (int(1) << (i * 2)) {
			return true
		}
	}
	return false
}

// https://www.cnblogs.com/grandyang/p/5403783.html
func solution2(num int) bool {
	if num > 0 && num & 0x55555555 == num && (num & (num - 1)) == 0 {
		return true
	} else {
		return false
	}
}

func isPowerOfFour(num int) bool {
	return solution2(num)
}
// @lc code=end

