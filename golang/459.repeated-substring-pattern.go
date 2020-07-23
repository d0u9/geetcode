/*
 * @lc app=leetcode id=459 lang=golang
 *
 * [459] Repeated Substring Pattern
 */

// @lc code=start
func CompSubStr(sublen int, slen int, s string) bool {
	fmt.Printf("sublen=%d, slen=%d\n", sublen, slen);
	if slen % sublen != 0 || sublen == slen {
		return false
	}

	for i := 1; i < slen / sublen; i++ {
		for j := 0; j < sublen; j++ {
			cur := i * sublen + j;
			if s[cur] != s[j] {
				return false
			}
		}
	}

	return true
}

func solution1(s string) bool {
	l := len(s)
	if l <= 1 {
		return false
	}
	sc := s[0]
	for j := 1; j < l; j++ {
		if s[j] == sc {
			if CompSubStr(j, l, s) {
				return true
			}
		}
	}

	return false
}

func solution2(s string) bool {
	l := len(s)
	if l <= 1 {
		return false
	}

	lasthead, starthead := 0, 0
	i, j := 1, 0
	found := false
	for ;i < l; i++ {
		if !found {
			if s[i] == s[0] {
				j = 1
				found = true
				starthead = i
				//fmt.Printf("match head: i = %d\n", i)
			}
		} else {
			if lasthead == 0 && s[i] == s[0] {
				//fmt.Printf("find head druing comparing, i = %d\n", i)
				lasthead = i
			}
			if s[i] != s[j] {
				//fmt.Printf("mismatch, i = %d, j = %d\n", i, j)
				found = false
				if lasthead > 0 {
					i = lasthead - 1
					//fmt.Printf("i reset to %d\n", i)
				}
				lasthead = 0
			}
			j++
		}
	}
	fmt.Println(found, i, j, starthead)
	if found && j % starthead == 0 {
		return true
	}

	return false
}

func repeatedSubstringPattern(s string) bool {
	return solution2(s)
}
// @lc code=end

