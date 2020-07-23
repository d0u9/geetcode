/*
 * @lc app=leetcode id=443 lang=golang
 *
 * [443] String Compression
 */

// @lc code=start
func compress(chars []byte) int {
	l := len(chars)
	if l < 2 {
		return l
	}

	i, j, cnt := 1, 0, 1
	c := chars[0]
	for i, j = 1, 0; i <= l; i++ {
		//fmt.Printf("c = %c\n", c)
		if i < l && c == chars[i] {
			cnt++
			continue
		}

		//fmt.Printf("char = %c, cnt = %d\n", c, cnt)
		//fmt.Printf("i = %d, j = %d\n", i , j)

		chars[j] = c
		j++
		if i < l {
			c = chars[i]
		}
		if cnt <= 1 {
			continue
		}

		cntStr := strconv.Itoa(cnt)
		//fmt.Printf("cntStr len = %d\n", len(cntStr))
		for k := 0; k < len(cntStr); k, j = k+1, j+1 {
			chars[j] = cntStr[k]
		}

		cnt = 1
	}
	return j
}
// @lc code=end

