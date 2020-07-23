/*
 * @lc app=leetcode id=389 lang=golang
 *
 * [389] Find the Difference
 */

// @lc code=start
func findTheDifference(s string, t string) byte {
	arr := make([]int, 26, 26)
	l1, l2 := len(s), len(t)
	for i := 0; i < l1; i++ {
		idx := int(s[i] - 'a')
		arr[idx]++
	}

	for i := 0; i < l2; i++ {
		idx := int(t[i] - 'a')
		arr[idx]--
	}

	for i := 0; i < 26; i++ {
		if arr[i] < 0 {
			return byte(i + 'a')
		}
	}
	return 'x'
}
// @lc code=end

