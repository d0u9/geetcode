/*
 * @lc app=leetcode id=566 lang=golang
 *
 * [566] Reshape the Matrix
 */

// @lc code=start
func matrixReshape(nums [][]int, r int, c int) [][]int {
	// The height and width of the given matrix is in range [1, 100].
	ro, co := len(nums), len(nums[0])
	if ro * co != r * c {
		return nums
	}

	ret := make([][]int, 0)
	p := 0
	cur := make([]int, c, c)
	for i := 0; i < ro; i++ {
		for j := 0; j < co; j++ {
			cur[p] = nums[i][j]
			p++
			if p % c == 0 {
				ret = append(ret, cur)
				cur = make([]int, c, c)
				p = 0
			}
		}
	}

	return ret
}
// @lc code=end

