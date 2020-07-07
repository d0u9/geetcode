/*
 * @lc app=leetcode id=263 lang=golang
 *
 * [263] Ugly Number
 */

// @lc code=start
func isUgly(num int) bool {
	if num <= 0{
		return false
	}
	pr := []int{ 2, 3, 5, }
	i, ret := 0, true
	for num > 1 {
		q, r := num / pr[i], num % pr[i]
		if r != 0 {
			i++
			if i > 2 {
				ret = false
				break
			}
		} else {
			num = q
		}
	}
	return ret
}
// @lc code=end

