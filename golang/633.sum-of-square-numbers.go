/*
 * @lc app=leetcode id=633 lang=golang
 *
 * [633] Sum of Square Numbers
 */

// @lc code=start
import "math"

func solution1(c int) bool {
	if c < 0 {
		return false
	}
	mp := make(map[int]bool)
	
	i := 0
	for {
		p := i * i
		if p > c {
			break
		}
		mp[p] = true
		t := c - p
		if _, ok := mp[t]; ok {
			return true
		}
		i++
	}
	return false
}

func perfect_square(t int) bool {
	l, r := 0, t
	for l <= r {
		mid := (l + r) / 2
		if mid * mid == t {
			return true
		} else if mid * mid > t {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return false
}

func solution2(c int) bool {
	term := int(math.Sqrt(float64(c)))
	//for i := 0; i <= c / 2; i++ {
	for i := 0; i <= term; i++ {
		t := c - i * i
		if perfect_square(t) {
			return true
		}
	}
	return false
}

func judgeSquareSum(c int) bool {
    return solution2(c)
}
// @lc code=end

