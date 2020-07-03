/*
 * @lc app=leetcode id=231 lang=golang
 *
 * [231] Power of Two
 */

// @lc code=start
func solution1(n int) bool {
	if n <= 0 {
		return false
	}
	dict := []int { 0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4, }
	c := 0
	for i := 0; i < int(unsafe.Sizeof(n)); i++ {
		c = c + dict[uint(n & 0xf)]
		n = n >> 4
	}
	if c > 1 {
		return false
	}
	return true
}

func solution2(n int) bool {
	if n <= 0 {
		return false
	}

	for q := n; q != 0; q = q / 2 {
		if q % 2 != 0 && q != 1 {
			return false
		}
	}
	return true
}

func isPowerOfTwo(n int) bool {
    return solution2(n)
}
// @lc code=end

