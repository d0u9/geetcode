/*
 * @lc app=leetcode id=453 lang=golang
 *
 * [453] Minimum Moves to Equal Array Elements
 */

// @lc code=start
func solution1(nums []int) int {
	l := len(nums)
	cnt := 0
	if l <= 1 {
		return 0
	}

	max, idmax := math.MinInt32, -1
	min := math.MaxInt32
	for i, v := range(nums) {
		if v > max {
			max = v
			idmax = i
		}

		if v < min {
			min = v
		}
	}

	fmt.Printf("max = %d, min = %d, max_idx=%d\n", max, min, idmax)

	diff := max - min
	for diff != 0 {
		tmp_max, tmp_min := math.MinInt32, math.MaxInt32
		tmp_idmax := -1
		for i := 0; i < l; i++ {
			if i == idmax {
				continue
			}

			v := nums[i] + diff
			if v > tmp_max {
				tmp_max = v
				tmp_idmax = i
			}

			if v < tmp_min {
				tmp_min = v
			}
			nums[i] = v
		}
		cnt = cnt + diff
		max, min, idmax = tmp_max, tmp_min, tmp_idmax
		diff = max - min
	}

	return cnt
}

/*
 * Each time, increase n-1 elements by 1 means the
 * sum of array increaes (n-1).
 * After m times move, the sum of array increases 
 * (n-1) * m.
 */
func solution2(nums []int) int {
	s, min := 0, math.MaxInt32
	for _, v := range(nums) {
		s = s + v
		if v < min {
			min = v
		}
	}

	return s - len(nums) * min
}

func minMoves(nums []int) int {
	return solution2(nums)
}
// @lc code=end

