/*
 * @lc app=leetcode id=383 lang=golang
 *
 * [383] Ransom Note
 */

// @lc code=start
func solution1(ransomNote string, magazine string) bool {
	l1, l2 := len(ransomNote), len(magazine)
	if l1 > l2 {
		return false
	}
	mp := make(map[byte]int)
	for i := 0; i < l1; i++ {
		b := ransomNote[i]
		if v, ok := mp[b]; ok {
			mp[b] = v + 1
		} else {
			mp[b] = 1
		}
	}

	for i := 0; i < l2; i ++ {
		b := magazine[i]
		if v, ok := mp[b]; ok {
			if v - 1 == 0 {
				delete(mp, b)
			} else {
				mp[b] = v - 1
			}
		}
	}

	return len(mp) == 0
}


func solution2(ransomNote string, magazine string) bool {
	return false
}

func canConstruct(ransomNote string, magazine string) bool {
	return solution1(ransomNote, magazine)
}
// @lc code=end

