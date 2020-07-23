/*
 * @lc app=leetcode id=405 lang=golang
 *
 * [405] Convert a Number to Hexadecimal
 */

// @lc code=start
func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	mp := []int{ '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f', }
	ret, l := []byte{}, 0
	for i := 7; i >= 0; i-- {
		v := (num >> (i*4)) & 0xf
		if v == 0 && l == 0 {
			continue
		}
		ret = append(ret, byte(mp[v]))
		l++
	}
	//fmt.Println(string(ret))
	return string(ret)
}
// @lc code=end

