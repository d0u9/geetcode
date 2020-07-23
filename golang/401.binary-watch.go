/*
 * @lc app=leetcode id=401 lang=golang
 *
 * [401] Binary Watch
 */

// @lc code=start

func countBits(num int) int {
	ret := 0
	for i := 0; i < 8; i++ {
		if ((num >> i) & (0x01)) == 0x01 {
			ret++
		}
	}
	return ret
}

func solution1(num int) []string {
	ret := []string{}
	for i := 0; i <= num; i++ {
		h, m := i, num - i
		if h > 4 || m > 8 {
			break
		}

		for j := 0; j < 12; j++ {
			if countBits(j) == h {
				for k := 0; k < 60; k++ {
					if countBits(k) == m {
						ret = append(ret, fmt.Sprintf("%d:%02d", j, k))
					}
				}
			}
		}
	}
	return ret
}

func solution2(num int) []string {
	
}

func readBinaryWatch(num int) []string {
	return solution1(num)
}
// @lc code=end

