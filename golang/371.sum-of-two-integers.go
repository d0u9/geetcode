/*
 * @lc app=leetcode id=371 lang=golang
 *
 * [371] Sum of Two Integers
 */

// @lc code=start
func getSum(a int, b int) int {
	w := int(unsafe.Sizeof(a)) * 4
	ret := int32(0)
	fmt.Printf("%b, %b\n", a, b)
	for i, c := 0, 0; i < w; i++ {
		ba, bb := (a >> i) & 0x1, (b >> i) & 0x1
		fmt.Printf("ba = %d, bb = %d\n", ba, bb)
		bm := ba ^ bb
		b := ba ^ bb ^ c
		c = (ba & bb) | (bm & c)
		fmt.Printf("b = %d, c = %d, bm = %d\n", b, c, bm)
		ret = ret | int32((b << i))
		fmt.Printf("%b\n", ret)
	}
	return int(ret)
}
// @lc code=end

