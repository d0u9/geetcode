/*
 * @lc app=leetcode id=598 lang=golang
 *
 * [598] Range Addition II
 */

// @lc code=start

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func solution1(m int, n int, ops [][]int) int {
	row, col := m, n
	max := 0
	for _, op := range(ops) {
		r, c := op[0], op[1]
		row, col = MinInt(r, row), MinInt(c, col)
		max++
	}

	return (col* row)
}

func maxCount(m int, n int, ops [][]int) int {
    return solution1(m, n, ops)
}
// @lc code=end

