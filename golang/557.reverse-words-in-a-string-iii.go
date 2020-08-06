/*
 * @lc app=leetcode id=557 lang=golang
 *
 * [557] Reverse Words in a String III
 */

// @lc code=start
func reverse(str []byte, l int) {
	//fmt.Println(string(str))
	for i := 0; i < l / 2; i ++ {
		str[i], str[l - i - 1] = str[l - i - 1], str[i]
	}
}

func reverseWords(s string) string {
	l := len(s)
	if l <= 0 {
		return ""
	}
	wl := 0
	str := []byte(s)
	for i := 0; i < l; i++ {
		if str[i] == ' ' && wl != 0 {
			reverse(str[ i - wl : i ], wl)
			wl = 0
		} else if str[i] != ' ' {
			wl++
		}
	}
	reverse(str[l - wl: l], wl)
	return string(str)
}
// @lc code=end

