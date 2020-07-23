/*
 * @lc app=leetcode id=476 lang=golang
 *
 * [476] Number Complement
 */

// @lc code=start
func findComplement(num int) int {
	if num == 0 {
		return 1
	}
	v := uint32(num)
	mask := uint32(0xffffffff)
	for w := 0; w < 32; w++ {
		if v & mask == 0 {
			break
		}
		mask = mask << 1
	}
	return int(v ^ (^mask))
}
// @lc code=end

