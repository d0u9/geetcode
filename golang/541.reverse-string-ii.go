/*
 * @lc app=leetcode id=541 lang=golang
 *
 * [541] Reverse String II
 */

// @lc code=start
func reverse(str []byte, l int) {
	//fmt.Println(string(str))
	for i := 0; i < l / 2; i++ {
		str[i], str[l - i - 1] = str[l - i - 1], str[i]
	}
	//fmt.Println(string(str))
}

func reverseStr(s string, k int) string {
	l := len(s)
	if l == 0 || k == 0 {
		return s
	}

	//fmt.Println(l, k)

	str := []byte(s)

	cof := 1
	for i := 0; i < l; i++ {
		if i == (cof * k) - 1 {
			if cof % 2 != 0 {
				reverse( str[ i - (k - 1) : i + 1 ], k )
			}
			cof++
		}
	}

	//fmt.Println(cof)

	if cof % 2 == 1 && l < k * cof {
		reverse(str[ (cof - 1) * k: ], l - (cof - 1) * k)
	}


	return string(str)
}
// @lc code=end

