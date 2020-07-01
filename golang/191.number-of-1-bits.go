/*
 * @lc app=leetcode id=191 lang=golang
 *
 * [191] Number of 1 Bits
 */

// @lc code=start
func hammingWeight(num uint32) int {
	table := [16]int {0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4,}
	mask, n := uint32(0xf), 0
	for i := 0; i < 8; i++ {
		n = n + table[num & mask]
		num = num >> 4
	}
	return n
}
// @lc code=end

