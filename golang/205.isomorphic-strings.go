/*
 * @lc app=leetcode id=205 lang=golang
 *
 * [205] Isomorphic Strings
 */

// @lc code=start
func slicestruct(s string, t string) bool {
	l := len(s)
	mp1 := make([]byte, 256, 256)
	mp2 := make([]byte, 256, 256)
	for i := 0; i < 256; i++ {
		mp1[i] = 0xff
		mp2[i] = 0xff
	}

	for i := 0; i < l; i++ {
		ids, idt := int(s[i]), int(t[i])
		if mp1[ids] == 0xff && mp2[idt] == 0xff {
			mp1[ids] = t[i]
			mp2[idt] = s[i]
		} else {
			if mp1[ids] != t[i] || mp2[idt] != s[i] {
				return false
			}
		}
	}
	return true
}

func mapstruct(s string, t string) bool {
	l := len(s)
	mp1 := make(map[byte]byte)
	mp2 := make(map[byte]byte)

	for i := 0; i < l; i++ {
		v1, ok1 := mp1[s[i]]
		v2, ok2 := mp2[t[i]]

		if !ok1 && !ok2 {
			mp1[s[i]] = t[i]
			mp2[t[i]] = s[i]
		} else {
			if !ok1 || !ok2 {
				return false
			}

			if v1 != t[i] || v2 != s[i] {
				return false
			}
		}
	}
	return true
}

func isIsomorphic(s string, t string) bool {
	return mapstruct(s, t)
}
// @lc code=end

