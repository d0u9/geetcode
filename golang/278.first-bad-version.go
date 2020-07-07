/*
 * @lc app=leetcode id=278 lang=golang
 *
 * [278] First Bad Version
 */

// @lc code=start
/** 
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad 
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	l, r := 1, n
	for l < r {
		c := (l + r) / 2
		if isBadVersion(c) {
			r = c - 1
		} else {
			l = c + 1
		}
	}

	if isBadVersion(l) {
		return l
	}
	return l + 1
}
// @lc code=end

