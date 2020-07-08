/*
 * @lc app=leetcode id=299 lang=golang
 *
 * [299] Bulls and Cows
 */

// @lc code=start
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func getHint(secret string, guess string) string {
	bull, l := 0, len(guess)
	mp1 := make([]int, 10, 10)
	mp2 := make([]int, 10, 10)

	for i := 0; i < l; i++ {
		id1, id2 := int(secret[i] - '0'), int(guess[i] - '0')
		if secret[i] == guess[i] {
			bull++
		}
		mp1[id1], mp2[id2] = mp1[id1] + 1, mp2[id2] + 1
	}
	
	cow := 0
	for i := 0; i < 10; i++ {
		cow += min(mp1[i], mp2[i])
	}

	cow = cow - bull
	return fmt.Sprintf("%dA%dB", bull, cow)
}
// @lc code=end

