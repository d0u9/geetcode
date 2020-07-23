/*
 * @lc app=leetcode id=482 lang=golang
 *
 * [482] License Key Formatting
 */

// @lc code=start

// ""5F3Z-2eqty-9-w"\n4"
// ""2-4A0r7-4k"\n3"
func cvt(b byte) byte {
	if b >= 'a' && b <= 'z' {
		return b - ('a' - 'A')
	}
	return b
}
func licenseKeyFormatting(S string, K int) string {
	l := len(S)
	cnt := 0
	ret := []byte{}
	
	for i := 0; i < l; i++ {
		if S[i] != '-' {
			cnt++
		}
	}

	fmt.Println(cnt % K)

	cnt = cnt % K
	i, j := 0, 0
	for ;cnt > 0; i++ {
		if S[i] == '-' {
			continue
		}
		ret = append(ret, cvt(S[i]))
		j++
		cnt--
	}

	cnt = K
	for ;i < l; i++ {
		if S[i] == '-' {
			continue
		}
		if cnt == K && j != 0 {
			ret = append(ret, '-')
		}
		ret = append(ret, cvt(S[i]))
		j++
		cnt--
		if cnt == 0 {
			cnt = K
		}
	}

	return string(ret)
}
// @lc code=end

