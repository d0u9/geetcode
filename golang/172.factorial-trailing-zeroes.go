/*
 * @lc app=leetcode id=172 lang=golang
 *
 * [172] Factorial Trailing Zeroes
 */

// @lc code=start
func count5(v int) int {
	c, n := 0, v
	for n % 5 == 0 {
		c++
		n = n / 5
	}
	return c
}

func solution1(n int) int {
	ret := 0
	for i := 5; i <= n; i = i + 5 {
		ret = ret + count5(i)
	}
	return ret
}

func solution2(n int) int {
	ret := 0
	for n > 0 {
		ret = ret + (n / 5)
		n = n / 5
	}
	return ret
}

func trailingZeroes(n int) int {
	return solution2(n)
}
// @lc code=end

