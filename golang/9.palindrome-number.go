/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */

// @lc code=start
func isPalindrome(x int) bool {
    if x < 0 {
		return false
	}

	o := x
	z := 0
	for x != 0 {
		rem := x % 10
		z = z * 10 + rem
		x = x / 10
	}
	
	if z == o {
		return true
	} else {
		return false
	}
}
// @lc code=end

