/*
 * @lc app=leetcode id=190 lang=golang
 *
 * [190] Reverse Bits
 */

// @lc code=start
func reverseBits(num uint32) uint32 {
	w := int(unsafe.Sizeof(num)) * 8
	var ret uint32 = 0
	mask := uint32(0x1)
	for i := 0; i < w; i++ {
		ret = ret << 1
		ret = (num & mask) | ret
		num = num >> 1
	}
    return ret
}
// @lc code=end

