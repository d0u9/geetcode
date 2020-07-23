/*
 * @lc app=leetcode id=374 lang=golang
 *
 * [374] Guess Number Higher or Lower
 */

// @lc code=start
/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	left, right := 1, n
	for left < right {
		c := (left + right) / 2
		v := guess(c)
		//fmt.Printf("v = %d, guess=%d\n", c, v)
		if v == -1 {
			right = c - 1
		} else if v == 1 {
			left = c + 1
		} else {
			return c
		}
	}

	//fmt.Println(left)

	if guess(left) == 0 {
		return left
	}
	return -1
}
// @lc code=end

