/*
 * @lc app=leetcode id=125 lang=golang
 *
 * [125] Valid Palindrome
 */

// @lc code=start
func testChar(c byte) (byte, bool) {
	if c >= 'a' && c <= 'z' {
		return c, true
	}

	if c >= 'A' && c <= 'Z' {
		return c + 32, true
	}

	if c >= '0' && c <= '9' {
		return c, true
	}

	return c, false
}


func isPalindrome(s string) bool {
	l := len(s)
	if l == 1 || l == 0 {
		return true
	}

	for i, j := 0, l - 1; i <= j; {
		var c1, c2 byte
		var ok bool
		if c1, ok = testChar(s[i]); !ok {
			i++
			continue
		}
		if c2, ok = testChar(s[j]); !ok {
			j--
			continue
		}
		//fmt.Println(c1, c2)
		if c1 != c2 {
			return false
		}
		i, j = i+1, j-1
	}
	return true
}
// @lc code=end

