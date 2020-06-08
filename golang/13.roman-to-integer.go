/*
 * @lc app=leetcode id=13 lang=golang
 *
 * [13] Roman to Integer
 */

// @lc code=start
func romanToInt(s string) int {
    dict := map[byte]int {
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	last := dict['M']
	result := 0
	for i := 0; i < len(s); i++ {
		cur := dict[s[i]]
		if cur > last {
			result = result - last
			result = result + (cur - last)
		} else {
			result = result + cur
		}
		last = cur
	}
	return result
}
// @lc code=end

