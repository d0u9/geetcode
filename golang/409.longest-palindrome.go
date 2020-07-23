/*
 * @lc app=leetcode id=409 lang=golang
 *
 * [409] Longest Palindrome
 */

// @lc code=start
func getidx(c byte) int {
	if c >= 'A' && c <= 'Z' {
		return int(c - 'A') + 26
	}
	return int(c - 'a')
}

func longestPalindrome(s string) int {
	l := len(s)
	mp := make([]int, 26 * 2, 26 * 2)
	//single_cnt, double_cnt := 0, 0
	for i := 0; i < l; i++ {
		c := s[i]
		idx := getidx(c)
		mp[idx] = mp[idx] + 1
	}

	ret, single := 0, 0
	for _, v := range(mp) {
		ret = ret + (int(v / 2) * 2)
		if v % 2 != 0 {
			single = single | 0x01
		}
	}

	return ret + single
}
// @lc code=end

