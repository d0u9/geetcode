/*
 * @lc app=leetcode id=520 lang=golang
 *
 * [520] Detect Capital
 */

// @lc code=start
func IsCapital(b byte) bool {
	if b >= 'A' && b <= 'Z' {
		return true
	}
	return false
}

func detectCapitalUse(word string) bool {
	l := len(word)
	if l <= 1 {
		return true
	}
	
	capital := IsCapital(word[1])
	if !IsCapital(word[0]) && capital {
		return false
	}
	
	for i := 2; i < l; i++ {
		if IsCapital(word[i]) != capital {
			return false
		}
	}

	return true
}
// @lc code=end

