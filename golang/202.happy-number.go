/*
 * @lc app=leetcode id=202 lang=golang
 *
 * [202] Happy Number
 */

// @lc code=start
func cal(n int) int {
	sum := 0
	for n != 0 {
		d := n % 10
		sum = sum + d * d
		n = n / 10
	}
	return sum
}

func isHappy(n int) bool {
	m := make(map[int]bool)
	for {
		if _, ok := m[n]; !ok {
			m[n] = true
			n = cal(n)
			if n == 1 {
				return true
			}
		} else {
			return false
		}
	}
	return false
}
// @lc code=end

