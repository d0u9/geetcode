/*
 * @lc app=leetcode id=575 lang=golang
 *
 * [575] Distribute Candies
 */

// @lc code=start
func solution1(candies []int) int {
	l := len(candies)
	mp := make(map[int]int)
	for i := 0; i < l; i++ {
		c := candies[i]
		mp[c]++
	}

	cnt := len(mp)
	if cnt >= (l / 2) {
		return l / 2
	}

	return cnt
}

func solution2(candies []int) int {
	l := len(candies)
	if l == 0 {
		return 0
	}
	sort.Ints(candies)

	cnt, last := 1, candies[0]
	for i := 1; i < l; i++ {
		if candies[i] != last {
			last = candies[i]
			cnt++
		}
	}

	if cnt >= (l / 2) {
		return l / 2
	}

	return cnt	
}

func solution3(candies []int) int {
	l := len(candies)
	if l == 0 {
		return 0
	}
	sort.Ints(candies)

	cnt, last := 1, candies[0]
	for i := 1; i < l; i++ {
		if candies[i] != last {
			last = candies[i]
			cnt++
		}
	}

	if cnt >= (l / 2) {
		return l / 2
	}

	return cnt	
}

func distributeCandies(candies []int) int {
	return solution2(candies)
}
// @lc code=end

