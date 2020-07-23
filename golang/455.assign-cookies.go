/*
 * @lc app=leetcode id=455 lang=golang
 *
 * [455] Assign Cookies
 */

// @lc code=start
func findContentChildren(g []int, s []int) int {
	cnt := 0
	l1, l2 := len(g), len(s)
	sort.Ints(g)
	sort.Ints(s)
	for i, j := 0, 0; i < l1 && j < l2; {
		if s[j] >= g[i] {
			cnt++
			i, j = i + 1, j + 1
		} else {
			j++
		}
	}
	return cnt
}
// @lc code=end

