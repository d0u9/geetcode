/*
 * @lc app=leetcode id=242 lang=golang
 *
 * [242] Valid Anagram
 */

// @lc code=start
func isAnagram(s string, t string) bool {
	l, l2 := len(s), len(t)
	if l != l2 {
		return false
	}
	a := make([]int, 26)
	for i := 0; i < l; i++ {
		id := int(s[i] - 'a')
		a[id] = a[id] + 1
	}
	for i := 0; i < l; i++ {
		id := int(t[i] - 'a')
		a[id] = a[id] - 1
	}

	for _, v := range a {
		if v != 0 {
			return false
		}
	}
	return true
}
// @lc code=end

