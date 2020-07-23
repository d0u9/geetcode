/*
 * @lc app=leetcode id=461 lang=golang
 *
 * [461] Hamming Distance
 */

// @lc code=start
func hammingDistance(x int, y int) int {
	xor := x ^ y
	cnt := 0
	for i := 0; i < 32; i++ {
		if v := xor & 0x01; v > 0 {
			cnt++
		}
		xor = xor >> 1
	}
    return cnt
}
// @lc code=end

