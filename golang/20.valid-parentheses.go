/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */

// @lc code=start
func isValid(s string) bool {
	l := len(s)
	if l == 0 {
		return true
	}

	if l % 2 != 0 {
		return false
	}

	m := map[rune]rune {
		'(': ')',
		'{': '}',
		'[': ']',
		')': '(',
		'}': '{',
		']': '[',
	}

	a := []rune {'X', }
	i := 0
	for _, c := range s {
		if m[c] != a[i] {
			a = append(a, c)
			i++
		} else {
			a = a[:i]
			i--
		}
	}

	if len(a) == 1 {
		return true
	}

	return false
}
// @lc code=end

