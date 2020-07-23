/*
 * @lc app=leetcode id=415 lang=golang
 *
 * [415] Add Strings
 */

// @lc code=start
func addStrings(num1 string, num2 string) string {
	c := int(0)
	ret := []byte{}
	i, j := len(num1) - 1, len(num2) - 1
	for ; i >= 0 || j >= 0; i, j = i - 1, j - 1 {
		var c1, c2 byte
		if i >= 0 && j >= 0 {
			c1, c2 = num1[i], num2[j]
		} else if i >= 0 {
			c1, c2 = num1[i], '0'
		} else {
			c1, c2 = '0', num2[j]
		}

		sum := int(c1 - '0') + int(c2 - '0') + c
		ret = append(ret, byte((sum % 10) + '0'))
		c = sum / 10
	}

	if c > 0 {
		ret = append(ret, '1')
	}

	l := len(ret)
	for i := 0; i < l / 2; i++ {
		ret[i], ret[l -i - 1] = ret[l - i - 1], ret[i]
	}

	return string(ret)
}
// @lc code=end

