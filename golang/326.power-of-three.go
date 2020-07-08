/*
 * @lc app=leetcode id=326 lang=golang
 *
 * [326] Power of Three
 */

// @lc code=start
func solution1(n int) bool {
    for n > 1 {
		r, q := n % 3, n / 3
		if r != 0 {
			return false
		}
		n = q
	}
	if n != 1 {
		return false
	}
	return true
}

// https://www.jianshu.com/p/f8feea78131b
func solution2(n int) bool {
	if n <= 0 {
		return false
	}
	if 1162261467 % n == 0 {
		return true
	} else {
		return false
	}
}


func isPowerOfThree(n int) bool {
	return solution1(n)
}

// @lc code=end

