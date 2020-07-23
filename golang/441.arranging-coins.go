/*
 * @lc app=leetcode id=441 lang=golang
 *
 * [441] Arranging Coins
 */

// @lc code=start
func solution1(n int) int {
	z := float64(8 * n + 1)
	k := (int(math.Sqrt(z)) - 1) / 2
	return int(k)
}

func cal(row int) int {
	return int((1 + row) * row / 2)
}

func solution2(n int) int {
	l, r := 0, n
	for l < r {
		c := (l + r) / 2
		v := cal(c)
		if v > n {
			r = c - 1
		} else if v < n {
			l = c + 1
		} else {
			return c
		}
	}

	if cal(l) > n {
		return l - 1
	}
	return l
}

func arrangeCoins(n int) int {
	return solution2(n)
}
// @lc code=end

