/*
 * @lc app=leetcode id=387 lang=golang
 *
 * [387] First Unique Character in a String
 */

// @lc code=start
func solution1(s string) int {
	arr := make([]int, 26, 26)
	l := len(s)
	for i := 0; i < l; i++ {
		c := int(s[i] - 'a')
		arr[c]++
	}

	for i := 0; i < l; i++ {
		c := int(s[i] - 'a')
		if v := arr[c]; v == 1 {
			return i
		}
	}
	return -1
}

func solution2(s string) int {
	return 0
}

func firstUniqChar(s string) int {
	return solution1(s)
}
// @lc code=end

