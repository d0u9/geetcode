/*
 * @lc app=leetcode id=119 lang=golang
 *
 * [119] Pascal's Triangle II
 */

// @lc code=start
func getRow(rowIndex int) []int {
	a := make([]int, rowIndex + 1)
	b := make([]int, rowIndex + 1)
	last := &a
	cur := &b

	(*last)[0] = 1
	for i := 1; i < rowIndex + 1; i++ {
		(*cur)[0], (*cur)[i] = 1, 1
		for j := 1; j < i; j++ {
			(*cur)[j] = (*last)[j - 1] + (*last)[j]
		}
		t := last
		last = cur
		cur = t
	}

	return (*last)
}
// @lc code=end

