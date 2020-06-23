/*
 * @lc app=leetcode id=118 lang=golang
 *
 * [118] Pascal's Triangle
 */

// @lc code=start
func generate(numRows int) [][]int {
    if numRows == 0 {
		return [][]int {}
	}

	ret := make([][]int, numRows)
	ret[0] = []int { 1, }
	last := ret[0]

	for i := 1; i < numRows; i++ {
		l := make([]int, i + 1)
		l[0], l[i] = 1, 1
		for j := 1; j < i; j++ {
			l[j] = last[j - 1] + last[j]
		}
		ret[i] = l
		last = l 
	}

	return ret
}
// @lc code=end

