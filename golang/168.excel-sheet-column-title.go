/*
 * @lc app=leetcode id=168 lang=golang
 *
 * [168] Excel Sheet Column Title
 */

// @lc code=start
func convertToTitle(n int) string {
	if n == 0 {
		return ""
	}
	ret := make([]byte, 0)
	l, v := 0, n -1
	for {
		q, r := int(v / 26), (v % 26)
		v = q - 1 
		ret = append(ret, byte('A' + r))
		l++
		if q == 0 {
			break
		}
	}
	//fmt.Println(string(ret))
	for i, j := 0, l - 1; i < l / 2; i, j = i+1, j-1 {
		ret[i], ret[j] = ret[j], ret[i]
	}
	return string(ret)
}
// @lc code=end

